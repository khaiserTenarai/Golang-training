package database

import (
    "context"
    "fmt"
    "net/url"

    "ecommerce_order_system/config"
    "github.com/jackc/pgx/v5"
)

func ConnectDB(cfg config.Config) (*pgx.Conn, error) {

    password := url.QueryEscape(cfg.DBPassword)

    connString := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s",
        cfg.DBUser,
        password,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
    )

    conn, err := pgx.Connect(context.Background(), connString)

    if err != nil {
        return nil, err
    }

    return conn, nil
}
