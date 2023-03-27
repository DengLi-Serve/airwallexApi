package airwallexApi

type apiInfo struct {
	method  string
	apiPath string
}

var apiPaths = map[string]map[string]map[string]apiInfo{
	"PaymentAcceptance": {
		"Customers": {
			"Create":               {MethodPost, "/api/v1/pa/customers/create"},
			"Retrieve":             {MethodGet, "/api/v1/pa/customers/%s"},
			"Update":               {MethodPost, "/api/v1/pa/customers/%s/update"},
			"GenerateClientSecret": {MethodGet, "/api/v1/pa/customers/%s/generate_client_secret"},
			"GetList":              {MethodGet, "/api/v1/pa/customers"},
		},
		"PaymentConsents": {
			"Create":   {MethodPost, "/api/v1/pa/payment_consents/create"},
			"Update":   {MethodPost, "/api/v1/pa/payment_consents/%s/update"},
			"Verify":   {MethodPost, "/api/v1/pa/payment_consents/%s/verify"},
			"Disable":  {MethodPost, "/api/v1/pa/payment_consents/%s/disable"},
			"Retrieve": {MethodGet, "/api/v1/pa/payment_consents/%s"},
			"GetList":  {MethodGet, "/api/v1/pa/payment_consents"},
		},
	},
}
