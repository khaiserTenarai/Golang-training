package database
import("context";"log";"q10-order-inventory-transaction/config";"github.com/jackc/pgx/v5/pgxpool")
func ConnectDB()*pgxpool.Pool{c:=config.Load();db,e:=pgxpool.New(context.Background(),c.DatabaseURL());if e!=nil{log.Fatal(e)};if e=db.Ping(context.Background());e!=nil{log.Fatal(e)};return db}
