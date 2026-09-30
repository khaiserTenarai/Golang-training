package database
import (
	"context"

	"attendance_leave/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {

	cfg := config.Load()

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
