package rave

type RaveClient struct {
	PaymentUrl    string            `json:"payment_url"`
	SecretKey     string            `json:"secret_key"`
	Customization RaveCustomization `json:"customizations"`
	Request       RaveRequest       `json:"request"`
	Response      RaveResponse      `json:"response"`
}

func NewRaveClient(secret_key string) *RaveClient {
	return &RaveClient{SecretKey: secret_key}
}

func NewCustomRaveClient(skey string, title string, description string, logo string) *RaveClient {
	raveClient := NewRaveClient(skey)
	raveClient.SetCustomization(title, description, logo)
	return raveClient
}

func (raveClient *RaveClient) SetSecretKey(skey string) {
	raveClient.SecretKey = skey
}

func (raveClient *RaveClient) SetCustomization(title string, description string, logo string) {
	raveClient.Customization = RaveCustomization{
		Title:       title,
		Description: description,
		Logo:        logo,
	}
}
