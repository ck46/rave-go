package rave

const (
	VerificationUrl = "https://api.flutterwave.com/v3/transactions/%s/verify"
	PaymentUrl      = "https://api.flutterwave.com/v3/payments"
)

type RaveCustomization struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type RaveCustomer struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Name        string `json:"name"`
}

type RaveRequest struct {
	TransactionRef string            `json:"tx_ref"`
	Amount         float64           `json:"amount"`
	Currency       string            `json:"currency"`
	PaymentOption  string            `json:"payment_options"`
	PaymentPlanID  string            `json:"payment_plan"`
	RedirectUrl    string            `json:"redirect_url"`
	Customer       RaveCustomer      `json:"customer"`
	Customization  RaveCustomization `json:"customizations"`
}

type RaveResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}
