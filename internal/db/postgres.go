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

func (P *PGConnector) SaveRates(ctx context.Context, rates []models.Rate) error {
	//TODO implement me
	panic("implement me")
}

func (P *PGConnector) GetRates(ctx context.Context) ([]models.Rate, error) {
	var rates = []models.Rate{}

	rows, err := P.PGConn.Query(
		ctx,
		`SELECT * FROM rates`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rate models.Rate
		err := rows.Scan(&rate)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}

	return rates, nil
}

func (P *PGConnector) GetCurrencyRate(
	ctx context.Context,
	rateReq models.CurrencyRate,
) (models.CurrencyRateResponse, error) {
	var curRatResp = models.CurrencyRateResponse{}
	rows, err := P.PGConn.Query(
		ctx,
		`SELECT * FROM rates
			WHERE name = $1 OR name = $2`,
		rateReq.ToCurrency, rateReq.FromCurrency)
	if err != nil {
		return curRatResp, err
	}

	var rates map[string]models.Rate
	for rows.Next() {
		var rate models.Rate
		err := rows.Scan(&rate)
		if err != nil {
			return curRatResp, err
		}
		rates[rate.Name] = rate
	}

	curRatResp.Rate = rates[rateReq.FromCurrency].Rate / rates[rateReq.ToCurrency].Rate

	return curRatResp, nil
}
