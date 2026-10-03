package database
import (
	"context"

	"ems/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

/*
	Database layer.

	Connect() creates a connection
	to PostgreSQL.
*/
func Connect() (*pgxpool.Pool, error) {

	// Load configuration.
	cfg := config.Load()

	/*
		Create PostgreSQL connection pool.
	*/
	db, err := pgxpool.New(
		context.Background(),
		cfg.DatabaseURL(),
	)

	if err != nil {
		return nil, err
	}

	/*
		Check whether PostgreSQL is available.
	*/
	err = db.Ping(context.Background())

	if err != nil {
		db.Close()

		return nil, err
	}

	return db, nil
}