package airwallex

type Option func(*airwallex)

func UpdateAuthorizationToken(status bool) Option {
	return func(a *airwallex) {
		a.autoUpdateAuthorizationToken = status
	}
}

func SetUrl(url string) Option {
	return func(a *airwallex) {
		a.url = url
	}
}

func SetFileUrl(fileUrl string) Option {
	return func(a *airwallex) {
		a.fileUrl = fileUrl
	}
}
