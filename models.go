package yookassa

type Amount struct {
	// Сумма в выбранной валюте. Выражается в виде строки и пишется через точку, например 10.00.
	// Количество знаков после точки зависит от выбранной валюты.
	Value string `json:"value"`
	// Код валюты в формате ISO-4217. Должен соответствовать валюте субаккаунта (recipient.gateway_id),
	// если вы разделяете потоки платежей, и валюте аккаунта (shopId в личном кабинете), если не разделяете.
	Currency string `json:"currency"`
}

type Card struct {
	Number        string `json:"number,omitempty"`
	CSC           string `json:"csc,omitempty"`
	Cardholder    string `json:"cardholder,omitempty"`
	First6        string `json:"first6,omitempty"`
	Last4         string `json:"last4,omitempty"`
	ExpiryMonth   string `json:"expiry_month,omitempty"`
	ExpiryYear    string `json:"expiry_year,omitempty"`
	CardType      string `json:"card_type,omitempty"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	IssuerName    string `json:"issuer_name,omitempty"`
}
