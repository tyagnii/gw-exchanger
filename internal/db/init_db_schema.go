package db

import (
	"context"
)

func (P *PGConnector) InitSchema(ctx context.Context) error {
	// todo: implement db schema init
	_, err := P.PGConn.Exec(
		ctx,
		`CREATE TABLE IF NOT EXISTS currencyRates (
    		id int Primary key,
    		name varchar UNIQUE NOT NULL,
    		rate float NOT NULL);`)
	if err != nil {
		return err
	}

	return nil
}
