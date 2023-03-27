package airwallex_struct

/*-PaymentAcceptance-*/

/*--Customers--*/

// PaymentAcceptanceCustomersCreateReq 创建客户请求参数(MerchantCustomerId, RequestId, Address.City, Address.CountryCode, Address.Street为必填参数)
type PaymentAcceptanceCustomersCreateReq struct {
	paymentAcceptanceCustomersInfo
}

//PaymentAcceptanceCustomersRetrieveReq 客户检索请求参数(Id 为必传参数)
type PaymentAcceptanceCustomersRetrieveReq struct {
	Id string `json:"-"` //客户ID
}

// PaymentAcceptanceCustomersUpdateReq 更新客户请求参数(Id, RequestId, Address.City, Address.CountryCode, Address.Street 为必传参数)
type PaymentAcceptanceCustomersUpdateReq struct {
	Id string `json:"-"` //客户ID
	paymentAcceptanceCustomersInfo
}

// PaymentAcceptanceCustomersGenerateClientSecretReq 生成客户秘钥请求参数(Id, RequestId, Address.City, Address.CountryCode, Address.Street 为必传参数)
type PaymentAcceptanceCustomersGenerateClientSecretReq struct {
	Id string `json:"-"` //客户ID
}

// PaymentAcceptanceCustomersGetListReq 获取客户列表请求参数
type PaymentAcceptanceCustomersGetListReq struct {
	FromCreatedAt      string `json:"from_created_at,omitempty"`      //ISO8601格式的开始时间
	MerchantCustomerId string `json:"merchant_customer_id,omitempty"` //商户端的客户ID
	PageNum            string `json:"page_num,omitempty"`             //页码从0开始
	PageSize           string `json:"page_size,omitempty"`            //每页要列出的客户数
	ToCreatedAt        string `json:"to_created_at,omitempty"`        //ISO8601格式的结束时间
}

/*--PaymentConsents--*/

// PaymentAcceptancePaymentConsentsCreateReq 创建付款授权书的请求参数(其中 CustomerId, NextTriggeredBy, RequestId 为必传参数)
type PaymentAcceptancePaymentConsentsCreateReq struct {
	ConnectedAccountId    string                 `json:"connected_account_id,omitempty"`    //Airwallex 分配关联实体的账户标识符（与平台关联的另一个账户,供平台使用,表示交易中的关联实体,平台是交易的所有者, 注意：这不能与 x-on-behalf-of 标头一起使用
	CustomerId            string                 `json:"customer_id,omitempty"`             //客户ID,将使用此付款授权书进行付款的客户的ID
	MerchantTriggerReason string                 `json:"merchant_trigger_reason,omitempty"` //指示后续付款是否已安排,仅在next_triggered_by=merchant时适用(scheduled|unscheduled; 默认：unscheduled)
	Metadata              map[string]interface{} `json:"metadata,omitempty"`                //附加的数据(json)
	NextTriggeredBy       string                 `json:"next_triggered_by,omitempty"`       //触发后续付款的一方,(merchant|customer)
	RequestId             string                 `json:"request_id,omitempty"`              //商户指定的唯一请求ID
}

// PaymentAcceptancePaymentConsentsUpdateReq 更新付款授权书的请求参数(其中 Id, RequestId 为必传参数)
type PaymentAcceptancePaymentConsentsUpdateReq struct {
	Id        string `json:"-"`                    //付款授权书ID
	Metadata  string `json:"metadata,omitempty"`   //附加数据(json)
	RequestId string `json:"request_id,omitempty"` //商户指定的唯一请求ID
}

// PaymentAcceptancePaymentConsentsVerifyReq 验证付款授权书的请求参数(其中 Id, RequestId 为必传参数)
type PaymentAcceptancePaymentConsentsVerifyReq struct {
	Id         string `json:"-"`
	Descriptor string `json:"descriptor,omitempty"` //验证期间可能向客户显示的描述符
	DeviceData struct {
		AcceptHeader string `json:"accept_header,omitempty"` //接受标头值
		Browser      struct {
			JavaEnabled       string `json:"java_enabled,omitempty"`
			JavascriptEnabled string `json:"javascript_enabled,omitempty"`
			UserAgent         string `json:"user_agent,omitempty"`
		} `json:"browser,omitempty"` //浏览器特定数据
		DeviceId         string   `json:"device_id,omitempty"`          //设备 ID 或广告 ID 或 IMEI
		IpAddress        string   `json:"ip_address,omitempty"`         //公网IP地址，同时支持IPv4和IPv6
		Language         string   `json:"language,omitempty"`           //语言或地区
		Location         struct{} `json:"location,omitempty"`           //位置数据
		Mobile           struct{} `json:"mobile,omitempty"`             //手机特定数据
		ScreenColorDepth int64    `json:"screen_color_depth,omitempty"` //以位为单位的屏幕颜色深度
		ScreenHeight     int64    `json:"screen_height,omitempty"`      //屏幕高度（以像素为单位）
		ScreenWidth      int64    `json:"screen_width,omitempty"`       //屏幕宽度（以像素为单位）
		Timezone         string   `json:"timezone,omitempty"`           //以小时为单位的 UTC 时区偏移量,必要时添加分钟数
	} `json:"device_data,omitempty"` //客户端设备信息,建议提供此数据以增加拥有无冲突 3DS 的几率
	PaymentMethod       string `json:"payment_method,omitempty"`
	RequestId           string `json:"request_id,omitempty"`
	ReturnUrl           string `json:"return_url,omitempty"`
	RiskControlOptions  string `json:"risk_control_options,omitempty"`
	VerificationOptions string `json:"verification_options,omitempty"`
}

type PaymentAcceptancePaymentConsentsDisableReq struct {
}

type PaymentAcceptancePaymentConsentsRetrieveReq struct {
}

type PaymentAcceptancePaymentConsentsGetListReq struct {
}
