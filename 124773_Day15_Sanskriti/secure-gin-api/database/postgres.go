package database

import (
 "context"
 "fmt"
 "time"
 "secure-gin-api/config"
 "github.com/jackc/pgx/v5/pgxpool"
)
func Connect(ctx context.Context, c config.Config) (*pgxpool.Pool,error) {
 dsn:=fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser,c.DBPassword,c.DBHost,c.DBPort,c.DBName)
 pc,err:=pgxpool.ParseConfig(dsn); if err!=nil{return nil,err}; pc.MaxConnIdleTime=5*time.Minute
 pool,err:=pgxpool.NewWithConfig(ctx,pc); if err!=nil{return nil,err}; if err=pool.Ping(ctx);err!=nil{pool.Close();return nil,err}; return pool,nil
}
