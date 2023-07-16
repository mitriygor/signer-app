package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"signer-api/config"
	"signer-api/pkg/db"
)

func main() {
	log.Println("Starting Signer API...")

	conn := db.ConnectToDB("DSN")
	connWriter := db.ConnectToDB("DSN_WRITER")

	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	if connWriter == nil {
		log.Panic("Can't connect to Postgres Writer!")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "",
		DB:       0,
	})

	defer conn.Close()
	defer connWriter.Close()

	router, err := initHandlers(conn, connWriter, redisClient)
	if err != nil {
		log.Panic(err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.WebPort),
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
