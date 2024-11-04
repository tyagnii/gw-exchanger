package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/tyagnii/gw-exchanger/internal/models"
)

type PGConnector struct {
	PGConn *pgx.Conn
	ctx    context.Context
}

func NewPGConnector(ctx context.Context, connectionString string) (*PGConnector, error) {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	return &PGConnector{PGConn: conn, ctx: ctx}, nil
}

func (P *PGConnector) SaveRates(ctx context.Context, rates models.Rates) error {
	//TODO implement me
	panic("implement me")
}

func (P *PGConnector) GetRates(ctx context.Context) ([]models.Rates, error) {
	//TODO implement me
	panic("implement me")
}

func (P *PGConnector) GetCurrencyRate(ctx context.Context, rate models.CurrencyRate) (models.CurrencyRateResponse, error) {
	//TODO implement me
	panic("implement me")
}
