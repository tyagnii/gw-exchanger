package db

import (
	"context"
	"github.com/tyagnii/gw-exchanger/internal/models"
)

type DBConnector interface {
	InitSchema(ctx context.Context) error
	SaveRates(ctx context.Context, rates []models.Rate) error
	GetRates(ctx context.Context) ([]models.Rate, error)
	GetCurrencyRate(ctx context.Context, rate models.CurrencyRate) (models.CurrencyRateResponse, error)
}
