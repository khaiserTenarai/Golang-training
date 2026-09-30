package database
import (
	"context"

	"salary-management/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg config.Config) (*pgxpool.Pool, error) {

	db, err := pgxpool.New(
		context.Background(),
		cfg.DatabaseURL(),
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
