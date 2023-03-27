package airwallex_struct

/*-PaymentAcceptance-*/

/*--Customers--*/

type paymentAcceptanceCustomersInfo struct {
	AdditionalInfo struct {
		FirstSuccessfulOrderDate string `json:"first_successful_order_date,omitempty"` //客户在商家网站上完成第一笔订单的日期
		RegisteredViaSocialMedia bool   `json:"registered_via_social_media,omitempty"` //客户是否通过社交媒体注册,默认为false
		RegistrationDate         string `json:"registration_date,omitempty"`           //客户在商家网站上注册的日期
	} `json:"additional_info"` //客户的附加信息

	Address struct {
		City        string `json:"city"`                   //地址->城市,最多 100 个字符
		CountryCode string `json:"country_code,omitempty"` //地址->国家代码,使用两个字符的 ISO标准 国家/地区代码
		Postcode    string `json:"postcode,omitempty"`     //地址->邮政编码,最多 10 个字符
		State       string `json:"state,omitempty"`        //地址->所在的州或省,最多 100 个字符
		Street      string `json:"street"`                 //地址->街道,最多 1000 个字符
	} `json:"address"` //客户地址

	BusinessName       string                 `json:"business_name,omitempty"` //客户的公司名称
	Email              string                 `json:"email,omitempty"`         //客户的电子邮件地址
	FirstName          string                 `json:"first_name,omitempty"`    //客户的名字
	LastName           string                 `json:"last_name,omitempty"`     //客户的姓氏
	MerchantCustomerId string                 `json:"merchant_customer_id"`    //该客户在商户系统中的唯一标识
	Metadata           map[string]interface{} `json:"metadata,omitempty"`      //元数据(您可以附加到该客户的信息json格式)
	PhoneNumber        string                 `json:"phone_number,omitempty"`  //客户的电话号码
	RequestId          string                 `json:"request_id"`              //商户指定的唯一请求ID
}

/*--PaymentConsents--*/

type paymentAcceptancePaymentConsentsInfo struct {
	ClientSecret           string `json:"client_secret,omitempty"`             //客户秘钥
	ConnectedAccountId     string `json:"connected_account_id,omitempty"`      //已连接账户的标识符
	CreatedAt              string `json:"created_at,omitempty"`                //创建时间
	CustomerId             string `json:"customer_id,omitempty"`               //客户ID
	Id                     string `json:"id,omitempty"`                        //此付款同意书的唯一标识符
	InitialPaymentIntentId string `json:"initial_payment_intent_id,omitempty"` //  使用此付款同意书确认的初始付款意向ID
	Mandate                struct {
		AcceptedAt      string `json:"accepted_at,omitempty"` //物者接受授权的时间
		BacsDirectDebit struct {
			DdiReference      string `json:"ddi_reference,omitempty"`       //直接借记工具的参考
			ServiceUserNumber string `json:"service_user_number,omitempty"` // 一个唯一的六位数号码，用于识别通过 Bacs 直接借记交易付款或收款的企业
		} `json:"bacs_direct_debit,omitempty"` //BACS直接借记授权信息
		Type    string `json:"type,omitempty"`    //授权书的付款方式类型(ach_direct_debit,bacs_direct_debit,becs_direct_debit)
		Version string `json:"version,omitempty"` //版本
	} `json:"mandate,omitempty"` //购物者对商家的授权，从其银行账户中扣款
	MerchantTriggerReason string                 `json:"merchant_trigger_reason,omitempty"` //仅在next_triggered_by=merchant适用(scheduled|unscheduled)
	Metadata              map[string]interface{} `json:"metadata,omitempty"`                //元数据(您可以附加到该客户的信息json格式)
	NextAction            struct {
		ContentType       string `json:"content_type,omitempty"`        //在method为POST时有值,如果没有值默认"application/json"
		Data              string `json:"data,omitempty"`                //数据(json)
		Email             string `json:"email,omitempty"`               //如果类型为notify_micro_deposits时,用于发送小额存款验证链接的电子邮件地址
		Method            string `json:"method,omitempty"`              //如果操作类型为redirect时的重定向方法 GET|POST
		MicroDepositCount int64  `json:"micro_deposit_count,omitempty"` //如果类型为notify_micro_deposits时,将存入购物者账户的资金数量
		Qrcode            string `json:"qrcode,omitempty"`              //有值标识你的操作类型是redirect,仅适用于微信
		Stage             string `json:"stage,omitempty"`               //请求所在的阶段
		Type              string `json:"type,omitempty"`                //下一次的操作类型(redirect|redirect_iframe|notify_micro_deposits)
		Url               string `json:"url,omitempty"`                 //重定向地址
	} `json:"next_action,omitempty"` // 商户的下一步行动
	NextTriggeredBy string `json:"next_triggered_by,omitempty"` //触发后续付款的一方(merchant|customer)
	PaymentMethod   struct {
		AchDirectDebit struct {
			AbaRoutingNumber string `json:"aba_routing_number,omitempty"` //用于识别美国银行的 9 位数字
			AccountNumber    string `json:"account_number,omitempty"`     //用于识别美国银行账户的 4-17 位数字
			BusinessAccount  bool   `json:"business_account,omitempty"`   //是否为商业账户
			OwnerEmail       string `json:"owner_email,omitempty"`        //账户持有人的电子邮件
			OwnerName        string `json:"owner_name,omitempty"`         //账户持有人姓名
		} `json:"ach_direct_debit,omitempty"` //当类型为ach_direct_debit,ACH的直接借记信息
		BacsDirectDebit struct {
			AccountNumber string `json:"account_number,omitempty"` //账户
			Address       struct {
				Line1    string `json:"line_1,omitempty"`   //门牌号码和街道名称
				Line2    string `json:"line_2,omitempty"`   // 地区名称
				Postcode string `json:"postcode,omitempty"` //地址的邮政编码
				Town     string `json:"town,omitempty"`     //镇
			} `json:"address,omitempty"` //账户地址信息
			BankName   string `json:"bank_name,omitempty"`   // 账户银行
			OwnerEmail string `json:"owner_email,omitempty"` //账户持有人的电子邮件
			OwnerName  string `json:"owner_name,omitempty"`  // 账户持有人姓名
			SortCode   string `json:"sort_code,omitempty"`   // 用于识别英国银行的 6 位数字
		} `json:"bacs_direct_debit,omitempty"` //当类型为ach_direct_debit,BACS的直接借记信息
		BecsDirectDebit struct {
			AccountNumber string `json:"account_number,omitempty"` //用于识别澳大利亚银行账户的 4-9 位数字
			BsbNumber     string `json:"bsb_number,omitempty"`     //6 位银行-州-分行编号
			OwnerEmail    string `json:"owner_email,omitempty"`    //账户持有人的电子邮件
			OwnerName     string `json:"owner_name,omitempty"`     //账户持有人姓名
		} `json:"becs_direct_debit,omitempty"` //当类型为becs_direct_debit,BECS的直接借记信息
		Card struct {
			AdditionalInfo struct {
				MerchantVerificationValue string `json:"merchant_verification_value,omitempty"` //商家验证值（由 VISA 在入职期间提供）或 MasterCard 分配的 ID（由 MasterCard 在入职期间提供）
				TokenRequestorId          string `json:"token_requestor_id,omitempty"`          //令牌请求者 ID（在入职期间由卡方案提供）
			} `json:"additional_info,omitempty"` //当 number_type 为 EXTERNAL_NETWORK_TOKEN时,外部网络令牌请求者的附加信息
			AvsCheck string `json:"avs_check,omitempty"` //地址是否通过校验
			Billing  struct {
				Address struct {
					City        string `json:"city,omitempty"`         // 地址的城市,最多 100 个字符
					CountryCode string `json:"country_code,omitempty"` //地址的国家代码,使用两个字符的 ISO 标准国家/地区代码
					Postcode    string `json:"postcode,omitempty"`     //地址的邮政编码,最多 10 个字符
					State       string `json:"state,omitempty"`        //地址所在的州或省,最多 100 个字符
					Street      string `json:"street,omitempty"`       //地址的街道,最多 1000 个字符
				} `json:"address,omitempty"` // 显示在信用卡发行人记录中的账单地
				DateOfBirth string `json:"date_of_birth,omitempty"` //客户的出日，格式为：YYYY-MM-DD
				Email       string `json:"email,omitempty"`         //客户的电子邮件地址
				FirstName   string `json:"first_name,omitempty"`    //客户的名字
				LastName    string `json:"last_name,omitempty"`     //客户的姓氏
				PhoneNumber string `json:"phone_number,omitempty"`  //客户的电话号码
			} `json:"billing,omitempty"` //账单信息
			Bin               string `json:"bin,omitempty"`                 //此卡的银行识别号码
			Brand             string `json:"brand,omitempty"`               //卡片品牌
			CardType          string `json:"card_type,omitempty"`           //卡的种类
			CvcCheck          string `json:"cvc_check,omitempty"`           //CVC是否通过校验
			ExpiryMonth       string `json:"expiry_month,omitempty"`        //代表卡到期月份的两位数字
			ExpiryYear        string `json:"expiry_year,omitempty"`         //代表卡到期年份的四位数字
			Fingerprint       string `json:"fingerprint,omitempty"`         //卡的指纹
			IsCommercial      bool   `json:"is_commercial,omitempty"`       //一个布尔值字段,表示该卡是否是商用的
			IssuerCountryCode string `json:"issuer_country_code,omitempty"` //发卡机构的国家代码
			IssuerName        string `json:"issuer_name,omitempty"`         //发行人名称
			Last4             string `json:"last_4,omitempty"`              //卡号后四位
			Name              string `json:"name,omitempty"`                //持卡人姓名
			NumberType        string `json:"number_type,omitempty"`         // 号码类型( PAN, EXTERNAL_NETWORK_TOKEN, AIRWALLEX_NETWORK_TOKEN)
		} `json:"card,omitempty"` //当类型为card,卡片信息
		Id   string `json:"id,omitempty"`   //为后续付款附加的付款方式的ID仅在type=card卡时提供
		Type string `json:"type,omitempty"` //付款方式(card, alipayhk, truemoney, gcash, dana, kakaopay, tng, wechatpay, ach_direct_debit, bacs_direct_debit, becs_direct_debit)
	} `json:"payment_method,omitempty"` //为后续付款附加的付款信息
	Purpose   string `json:"purpose,omitempty"`    //付款同意书的目的(one_off,recurring)
	RequestId string `json:"request_id,omitempty"` //商户指定的唯一请求ID
	Status    string `json:"status,omitempty"`     // PENDING_VERIFICATION (待验证:准备验证或验证过程中是否next_action返回以指导下一步) PENDING (待办:已提交付款授权书。 它仅在付款方式类型为 bacs_direct_debit 时发生, 您需要等待最终结果,如果授权被拒绝，状态将变为禁用;如果授权被接受,则状态将变为已验证) VERIFIED (已验证:付款授权书已通过验证，可用于后续付款) DISABLED (付款授权书被禁用,您无法再次验证它)
	UpdatedAt string `json:"updated_at,omitempty"` //上次更新的时间
}
