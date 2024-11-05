package server

import (
	"context"
	"github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"github.com/tyagnii/gw-exchanger/internal/db"
	"github.com/tyagnii/gw-exchanger/internal/models"

	"go.uber.org/zap"
)

type ExchangeServer struct {
	exchanger.UnimplementedExchangeServiceServer
	DBConn        db.DBConnector
	ServerAddress string
	Logger        *zap.SugaredLogger
}

func NewExchangeServer(DBConn db.DBConnector, ServerAddress string, Logger *zap.SugaredLogger) *ExchangeServer {
	return &ExchangeServer{
		DBConn:        DBConn,
		ServerAddress: ServerAddress,
		Logger:        Logger,
	}
}

func (s *ExchangeServer) GetExchangeRates(
	ctx context.Context,
	eEmpty *exchanger.Empty,
) (*exchanger.ExchangeRatesResponse, error) {
	var exResponse exchanger.ExchangeRatesResponse

	r, err := s.DBConn.GetRates(ctx)
	s.Logger.Debugf("GetExchangeRates s.DBConn.GetRates: %v", r)
	if err != nil {
		s.Logger.Error("error due DB request GetRates", err)
		return nil, err
	}

	for _, rr := range r {
		exResponse.Rates[rr.Name] = rr.Rate
	}
	s.Logger.Debugf("GetExchangeRates exResponse: %v", exResponse)

	return &exResponse, nil
}

func (s *ExchangeServer) GetExchangeRateForCurrency(
	ctx context.Context,
	eCurReq *exchanger.CurrencyRequest,
) (*exchanger.ExchangeRateResponse, error) {
	var exResponse exchanger.ExchangeRateResponse
	var mCurRate models.CurrencyRate

	mCurRate.FromCurrency = eCurReq.FromCurrency
	mCurRate.ToCurrency = eCurReq.ToCurrency
	s.Logger.Debugf("GetExchangeRateForCurrency mCurRate: %v", mCurRate)

	r, err := s.DBConn.GetCurrencyRate(ctx, mCurRate)
	if err != nil {
		s.Logger.Error("error due DB request GetCurrencyRate", err)
		return nil, err
	}

	exResponse.Rate = r.Rate
	exResponse.ToCurrency = eCurReq.ToCurrency
	exResponse.FromCurrency = eCurReq.FromCurrency

	return &exResponse, nil
}
