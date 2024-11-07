package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/tyagnii/gw-exchanger/internal/models"
	"go.uber.org/zap"
	"time"
)

type PGConnector struct {
	PGConn *pgx.Conn
	ctx    context.Context
}

func NewPGConnector(ctx context.Context, connectionString string, sLogger *zap.SugaredLogger) (*PGConnector, error) {
	// Create db connection
	var DBconn *pgx.Conn
	go func(ctx context.Context) {
		var err error
		for {
			DBconn, err = pgx.Connect(ctx, connectionString)
			if err != nil {
				sLogger.Errorf("Error connecting to database: %v", err)
			} else {
				sLogger.Debugf("Connected to database. Breaking loop")
				break
			}

			switch ctx.Err() {
			case context.Canceled:
				break
			case context.DeadlineExceeded:
				break
			default:
				sLogger.Error("Could not connect to database")
				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	<-ctx.Done()
	if DBconn == nil {
		return nil, fmt.Errorf("Timeout connecting to database")
	}

	// Return PGConnector with new context
	return &PGConnector{PGConn: DBconn, ctx: context.Background()}, nil
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
