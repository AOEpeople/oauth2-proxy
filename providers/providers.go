package providers

import (
	"github.com/bitly/go-simplejson"
)

type Provider interface {
	Data() *ProviderData
	GetEmailAddress(auth_response *simplejson.Json,
		access_token string) (string, error)
}

func New(provider string, p *ProviderData) Provider {
	switch provider {
	default:
		return NewGoogleProvider(p)
	}
}
