package airwallex

import (
	"airwallex/airwallex_struct"
	"fmt"
)

func (a *airwallex) obtainAccessToken() error {
	a.request = &request{
		method:  MethodPost,
		heard:   map[string]string{"Authorization": "", "x-api-key": a.apiKey, "x-client-id": a.clientId},
		param:   nil,
		apiPath: "/api/v1/authentication/login",
	}
	type Resp struct {
		ExpiresAt string `json:"expires_at"`
		Token     string `json:"token"`
	}
	resp := &Resp{}
	err := a.sendRequest(&resp)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	if err != nil {
		return err
	}
	AuthorizationToken = resp.Token
	return nil
}

/*-PaymentAcceptance-*/

/*--Customers--*/

// PaymentAcceptanceCustomersCreate 创建客户 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Customers/_api_v1_pa_customers_create/post
func (a *airwallex) PaymentAcceptanceCustomersCreate(param airwallex_struct.PaymentAcceptanceCustomersCreateReq, resp *airwallex_struct.PaymentAcceptanceCustomersCreateResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["Customers"]["Create"].method,
		apiPath: apiPaths["PaymentAcceptance"]["Customers"]["Create"].apiPath,
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptanceCustomersRetrieve 客户检索 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Customers/_api_v1_pa_customers__id_/get
func (a *airwallex) PaymentAcceptanceCustomersRetrieve(param airwallex_struct.PaymentAcceptanceCustomersRetrieveReq, resp *airwallex_struct.PaymentAcceptanceCustomersRetrieveResp) error {
	a.request = &request{
		method:  apiPaths["PaymentAcceptance"]["Customers"]["Retrieve"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["Customers"]["Retrieve"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptanceCustomersUpdate 修改客户信息 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Customers/_api_v1_pa_customers__id__update/post
func (a *airwallex) PaymentAcceptanceCustomersUpdate(param airwallex_struct.PaymentAcceptanceCustomersUpdateReq, resp *airwallex_struct.PaymentAcceptanceCustomersUpdateResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["Customers"]["Update"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["Customers"]["Update"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptanceCustomersGenerateClientSecret 生成客户秘钥 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Customers/_api_v1_pa_customers__id__generate_client_secret/get
func (a *airwallex) PaymentAcceptanceCustomersGenerateClientSecret(param airwallex_struct.PaymentAcceptanceCustomersGenerateClientSecretReq, resp *airwallex_struct.PaymentAcceptanceCustomersGenerateClientSecretResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["Customers"]["GenerateClientSecret"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["Customers"]["GenerateClientSecret"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptanceCustomersGetList 获取客户列表 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Customers/_api_v1_pa_customers/get
func (a *airwallex) PaymentAcceptanceCustomersGetList(param airwallex_struct.PaymentAcceptanceCustomersGetListReq, resp *airwallex_struct.PaymentAcceptanceCustomersGetListResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["Customers"]["GetList"].method,
		apiPath: apiPaths["PaymentAcceptance"]["Customers"]["GetList"].apiPath,
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

/*--PaymentConsents--*/

// PaymentAcceptancePaymentConsentsCreate 创建付款同意书 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Payment_Consents/_api_v1_pa_payment_consents_create/post
func (a *airwallex) PaymentAcceptancePaymentConsentsCreate(param airwallex_struct.PaymentAcceptancePaymentConsentsCreateReq, resp *airwallex_struct.PaymentAcceptancePaymentConsentsCreateResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["PaymentConsents"]["Create"].method,
		apiPath: apiPaths["PaymentAcceptance"]["PaymentConsents"]["Create"].apiPath,
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptancePaymentConsentsUpdate 修改付款同意书 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Payment_Consents/_api_v1_pa_payment_consents__id__update/post
func (a *airwallex) PaymentAcceptancePaymentConsentsUpdate(param airwallex_struct.PaymentAcceptanceCustomersUpdateReq, resp *airwallex_struct.PaymentAcceptanceCustomersUpdateResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["PaymentConsents"]["Update"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["PaymentConsents"]["Update"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptancePaymentConsentsVerify 验证付款同意书 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Payment_Consents/_api_v1_pa_payment_consents__id__verify/post
func (a *airwallex) PaymentAcceptancePaymentConsentsVerify(param airwallex_struct.PaymentAcceptanceCustomersGenerateClientSecretReq, resp *airwallex_struct.PaymentAcceptanceCustomersGenerateClientSecretResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["PaymentConsents"]["Verify"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["PaymentConsents"]["Verify"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptancePaymentConsentsDisable 禁用付款同意书 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Payment_Consents/_api_v1_pa_payment_consents__id__disable/post
func (a *airwallex) PaymentAcceptancePaymentConsentsDisable(param airwallex_struct.PaymentAcceptanceCustomersGenerateClientSecretReq, resp *airwallex_struct.PaymentAcceptanceCustomersGenerateClientSecretResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["PaymentConsents"]["Disable"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["PaymentConsents"]["Disable"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptancePaymentConsentsRetrieve 检索付款同意书 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Payment_Consents/_api_v1_pa_payment_consents__id_/get
func (a *airwallex) PaymentAcceptancePaymentConsentsRetrieve(param airwallex_struct.PaymentAcceptanceCustomersRetrieveReq, resp *airwallex_struct.PaymentAcceptanceCustomersRetrieveResp) error {
	a.request = &request{
		method:  apiPaths["PaymentAcceptance"]["PaymentConsents"]["Retrieve"].method,
		apiPath: fmt.Sprintf(apiPaths["PaymentAcceptance"]["PaymentConsents"]["Retrieve"].apiPath, param.Id),
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}

// PaymentAcceptancePaymentConsentsGetList 获取付款同意书列表 对应官网参照地址:https://www.airwallex.com/docs/api#/Payment_Acceptance/Payment_Consents/_api_v1_pa_payment_consents/get
func (a *airwallex) PaymentAcceptancePaymentConsentsGetList(param airwallex_struct.PaymentAcceptanceCustomersGetListReq, resp *airwallex_struct.PaymentAcceptanceCustomersGetListResp) error {
	a.request = &request{
		param:   param,
		method:  apiPaths["PaymentAcceptance"]["PaymentConsents"]["GetList"].method,
		apiPath: apiPaths["PaymentAcceptance"]["PaymentConsents"]["GetList"].apiPath,
	}
	err := a.sendRequest(&resp)
	if err != nil {
		return err
	}
	return nil
}
