package models

type Rate struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

type CurrencyRate struct {
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
}

type CurrencyRateResponse struct {
	Rate float64 `json:"rate"`
}
