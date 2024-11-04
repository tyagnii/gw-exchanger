package models

type Rates struct {
	USD float64 `json:"USD"`
	RUB float64 `json:"RUB"`
	EUR float64 `json:"EUR"`
}

type CurrencyRate struct {
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
}

type CurrencyRateResponse struct {
	Rate float32 `json:"rate"`
}
