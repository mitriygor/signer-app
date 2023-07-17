package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"key-keeper-service/config"
	"key-keeper-service/pkg/db"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Signer API...")

	conn := db.ConnectToDB("DSN")

	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "",
		DB:       0,
	})

	defer conn.Close()

	router, err := initHandlers(conn, redisClient)
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
