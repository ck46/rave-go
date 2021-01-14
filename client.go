package rave

var payment_url = "https://api.flutterwave.com/v3/payments"

type RaveClient struct {
	PaymentUrl    string            `json:"payment_url"`
	SecretKey     string            `json:"secret_key"`
	Customization RaveCustomization `json:"customizations"`
	Request       RaveRequest       `json:"request"`
	Response      RaveResponse      `json:"response"`
}

func NewRaveClient() *RaveClient {
	return &RaveClient{PaymentUrl: payment_url}
}

func NewRaveClientWithSecretKey(skey string) *RaveClient {
	return &RaveClient{SecretKey: skey, PaymentUrl: payment_url}
}

func NewCustomRaveClient(skey string, title string, description string, logo string) *RaveClient {
	raveClient := NewRaveClientWithSecretKey(skey)
	raveClient.SetCustomization(title, description, logo)
	return raveClient
}

func (raveClient *RaveClient) SetPaymentUrl(payment_url string) {
	raveClient.PaymentUrl = payment_url
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
