package airwallex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	MethodPost  = "POST"
	MethodGet   = "GET"
	ContentType = "application/json"
	Url         = "https://api.airwallex.com"   //正式线
	FileUrl     = "https://files.airwallex.com" //正式线
)

type Airwallex struct {
	clientId                     string
	apiKey                       string
	url                          string
	fileUrl                      string
	request                      *request
	autoUpdateAuthorizationToken bool //是否自动刷新token(true:每25分钟刷新一次token,官方有效期为30分钟)
}

var AuthorizationToken string

func New(clientId, apiKey string, Option ...Option) *Airwallex {
	a := &Airwallex{
		clientId: clientId, //airwallex的clientId
		apiKey:   apiKey,   //airwallex的apiKey
		url:      Url,      //airwallex请求地址
		fileUrl:  FileUrl,  //airwallex文件请求地址
	}
	for _, opt := range Option {
		opt(a)
	}
	_ = a.obtainAccessToken()
	if a.autoUpdateAuthorizationToken {
		go func() {
			ticker := time.NewTicker(2 * time.Second)
			for {
				select {
				case <-ticker.C:
					_ = a.obtainAccessToken()
					fmt.Println(AuthorizationToken)
				}
			}
		}()
	}
	return a
}

type request struct {
	method    string            //发送请求的方式默认Post
	heard     map[string]string //请求头默认设置 {"Content-Type": "application/json","Authorization":全局token}
	param     interface{}       //请求所需要的参数
	apiPath   string            //要请求的api路径
	isFileUrl bool              //是否使用文件URL发送请求
}

func (a *Airwallex) sendRequest(resp interface{}) error {
	paramJson, _ := json.Marshal(a.request.param)
	param := bytes.NewBuffer(paramJson)
	client := &http.Client{}
	var url string
	if a.request.isFileUrl {
		url = a.fileUrl + a.request.apiPath
	} else {
		url = a.url + a.request.apiPath
	}
	if a.request.method == "" {
		a.request.method = MethodPost
	}
	fmt.Printf("request(method:%s,  url:%s,  param:%s)\n", a.request.method, url, param)
	req, err := http.NewRequest(a.request.method, url, param)
	if err != nil {
		return err
	}
	//header
	for k, v := range a.request.heard {
		if v == "" {
			continue
		}
		req.Header.Add(k, v)
	}
	if _, ok := a.request.heard["Content-Type"]; !ok {
		req.Header.Add("Content-Type", ContentType)
	}
	if _, ok := a.request.heard["Authorization"]; !ok {
		req.Header.Add("Authorization", AuthorizationToken)
	}

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	//可控范围内错误处理
	err = a.errorHandling(res)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(res.Body)
	//返回结果
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("request-err (%v)\n", err)
		return err
	}
	fmt.Printf("request-body(%v)\n", string(body))
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
}
