package main

import (
 "context"
 "database/sql"
 "fmt"
 "log"
 "net/http"
 "os"
 "os/signal"
 "syscall"
 "time"

 "employee-management/config"
 "employee-management/controller"
 "employee-management/middleware"
 "employee-management/repository"
 "employee-management/service"
 "github.com/gin-gonic/gin"
 _ "github.com/jackc/pgx/v5/stdlib"
 "github.com/joho/godotenv"
)

func main(){
 _=godotenv.Load(".env")
 cfg,err:=config.Load();if err!=nil{log.Fatal(err)}
 dsn:=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",cfg.DBHost,cfg.DBPort,cfg.DBUser,cfg.DBPassword,cfg.DBName)
 db,err:=sql.Open("pgx",dsn);if err!=nil{log.Fatal(err)};defer db.Close()
 db.SetMaxOpenConns(20);db.SetMaxIdleConns(5);db.SetConnMaxLifetime(30*time.Minute)
 ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second);defer cancel()
 if err=db.PingContext(ctx);err!=nil{log.Fatalf("PostgreSQL connection failed: %v",err)}

 // Dependency injection: repository -> service -> controller.
 repo:=repository.New(db);svc:=service.New(repo);ctrl:=controller.New(svc,db)
 gin.SetMode(gin.ReleaseMode);r:=gin.New();r.Use(gin.Recovery(),middleware.RequestID(),middleware.RequestLogger());ctrl.Routes(r)
 server:=&http.Server{Addr:":"+cfg.ServerPort,Handler:r,ReadHeaderTimeout:5*time.Second}
 go func(){log.Printf("API listening on :%s",cfg.ServerPort);if e:=server.ListenAndServe();e!=nil&&e!=http.ErrServerClosed{log.Fatalf("server: %v",e)}}()
 stop:=make(chan os.Signal,1);signal.Notify(stop,syscall.SIGINT,syscall.SIGTERM);<-stop
 shutdown,cancelShutdown:=context.WithTimeout(context.Background(),5*time.Second);defer cancelShutdown();if err:=server.Shutdown(shutdown);err!=nil{log.Printf("shutdown error: %v",err)}
}
