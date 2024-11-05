package models

type Rate struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float32 `json:"rate"`
}

type CurrencyRate struct {
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
}

type CurrencyRateResponse struct {
	Rate float32 `json:"rate"`
}
