package airwallexApi

type Option func(*Airwallex)

func UpdateAuthorizationToken(status bool) Option {
	return func(a *Airwallex) {
		a.autoUpdateAuthorizationToken = status
	}
}

func SetUrl(url string) Option {
	return func(a *Airwallex) {
		a.url = url
	}
}

func SetFileUrl(fileUrl string) Option {
	return func(a *Airwallex) {
		a.fileUrl = fileUrl
	}
}
