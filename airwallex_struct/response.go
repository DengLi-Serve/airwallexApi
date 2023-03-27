package airwallex_struct

// PaymentAcceptanceCustomersCreateResp 创建客户返回参数
type PaymentAcceptanceCustomersCreateResp struct {
	paymentAcceptanceCustomersInfo
	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
	UpdatedAt string `json:"updated_at"`
}

//PaymentAcceptanceCustomersRetrieveResp 客户检索返回参数
type PaymentAcceptanceCustomersRetrieveResp struct {
	paymentAcceptanceCustomersInfo
}

//PaymentAcceptanceCustomersUpdateResp 更新客户返回参数
type PaymentAcceptanceCustomersUpdateResp struct {
	paymentAcceptanceCustomersInfo
}

type PaymentAcceptanceCustomersGenerateClientSecretResp struct {
	ClientSecret string `json:"client_secret"` //客户秘钥
	ExpiredTime  string `json:"expired_time"`  //过期时间
}

// PaymentAcceptanceCustomersGetListResp 获取客户列表返回参数
type PaymentAcceptanceCustomersGetListResp struct {
	HasMore bool                                    `json:"has_more"` // 标识是否有更多的结果
	Items   []*PaymentAcceptanceCustomersCreateResp `json:"items"`    //客户信息数组
}

/*--PaymentConsents--*/

// PaymentAcceptancePaymentConsentsCreateResp 创建付款授权书返回参数
type PaymentAcceptancePaymentConsentsCreateResp struct {
	paymentAcceptancePaymentConsentsInfo
}

// PaymentAcceptancePaymentConsentsUpdateResp 更新付款授权书返回参数
type PaymentAcceptancePaymentConsentsUpdateResp struct {
	paymentAcceptancePaymentConsentsInfo
}

type PaymentAcceptancePaymentConsentsVerifyResp struct {
}

type PaymentAcceptancePaymentConsentsDisableResp struct {
}

type PaymentAcceptancePaymentConsentsRetrieveResp struct {
}

type PaymentAcceptancePaymentConsentsGetListResp struct {
}
