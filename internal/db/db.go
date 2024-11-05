package db

import (
	"context"
	"github.com/tyagnii/gw-exchanger/internal/models"
)

type DBConnector interface {
	InitSchema(ctx context.Context) error
	SaveRates(ctx context.Context, rates models.Rates) error
	GetRates(ctx context.Context) ([]models.Rates, error)
	GetCurrencyRate(ctx context.Context, rate models.CurrencyRate) (models.CurrencyRateResponse, error)
}
