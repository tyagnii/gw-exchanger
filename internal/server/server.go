package server

import (
	"context"
	"github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"github.com/tyagnii/gw-exchanger/internal/db"

	"go.uber.org/zap"
)

type ExchangeServer struct {
	exchanger.UnimplementedExchangeServiceServer
	DBConn        db.DBConnector
	ServerAddress string
	Logger        zap.SugaredLogger
}

func NewExchangeServer(DBConn db.DBConnector, ServerAddress string) *ExchangeServer {
	return &ExchangeServer{
		DBConn:        DBConn,
		ServerAddress: ServerAddress,
		Logger:        zap.SugaredLogger{},
	}
}

func (s *ExchangeServer) GetExchangeRates(
	context.Context,
	*exchanger.Empty,
) (*exchanger.ExchangeRatesResponse, error) {

	return nil, nil
}

func (s *ExchangeServer) GetExchangeRateForCurrency(
	context.Context,
	*exchanger.CurrencyRequest,
) (*exchanger.ExchangeRateResponse, error) {

	return nil, nil
}
