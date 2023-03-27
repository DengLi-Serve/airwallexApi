package airwallexApi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//http请求状态码
const (
	HttpStatusOk           = 200
	HttpStatusCreated      = 201
	HttpStatusBadRequest   = 400
	HttpStatusUnauthorized = 401
	HttpStatusNotFound     = 404
	HttpStatusServerError  = 500
)

var apiErrorCodes = map[string]string{
	"already_exists":                    "要求创建的实体已经存在",
	"amount_above_limit":                "请求的金额超过了最高支付或转换限额",
	"amount_above_payment_method_limit": "请求的金额超过了该支付方式和国家的最高限额",
	"amount_below_limit":                "请求的金额低于最低支付或转换限额",
	"can_not_be_edited":                 "付款正在处理中，因此不能再进行编辑",
	"conversion_create_failed":          "由于未知的原因，例如由于第三方的错误，该转换没有被创建",
	"credentials_expired":               "授权证书已过期，应申请新的授权证书",
	"credentials_invalid":               "授权证书无效或已失效；应申请新的授权证书",
	"field_required":                    "请求中缺少一个强制性字段",
	"invalid_argument":                  "提供的参数类型不正确或未通过验证测试",
	"invalid_currency_pair":             "货币对无效；例如，当买入货币等于卖出货币的时候",
	"invalid_payment_date":              "不能选择付款日期，因为它至少是在其中一个国家的假日/非营业日",
	"invalid_conversion_date":           "不能选择转换日期，因为它是在至少一个相关国家的假日/非营业日",
	"service_unavailable":               "由于工作时间的限制，或由于第三方服务暂时不可用，所要求的服务目前不可用",
	"term_agreement_is_required":        "需要同意服务条款，并且必须以真布尔值的形式提供",
	"too_many_requests":                 "收到太多的请求，速度太快",
	"unsupported_country_code":          "目前不支持向所请求的国家支付（或从该国支付）",
	"unsupported_currency":              "目前不支持从所请求的货币转换（或从货币转换）",
	"unsupported_payment_method":        "所选择的支付方式对所选择的受益银行国家和支付货币不可用",
}

var tokenTryNum int

func (a *Airwallex) errorHandling(res *http.Response) error {
	switch res.StatusCode {
	case HttpStatusOk, HttpStatusCreated:
		return nil
	case HttpStatusUnauthorized:
		if tokenTryNum == 4 {
			return errors.New("unauthorized try four error")
		}
		var e error
		e = a.obtainAccessToken()
		if e == nil {
			tokenTryNum = 0
			return nil
		}
		tokenTryNum += 1
		return e
	default:
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(fmt.Sprintf("HttpStatus:(%d)), Body(%s)", res.StatusCode, body))
	}
}
