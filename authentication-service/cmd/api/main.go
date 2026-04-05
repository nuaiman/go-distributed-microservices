package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const webPort = "8080"

type Application struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication Service")

	db := connectToDB()
	if db == nil {
		log.Fatal("Couldnot connect to Postgres")
	}

	app := Application{
		DB:     db,
		Models: data.New(db),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Printf("Starting authentication-service on port :%s\n", webPort)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for range 10 {
		db, err := sql.Open("pgx", dsn)
		if err == nil && db.Ping() == nil {
			log.Println("Couldnot cobbect to Postgres")
			return db
		}

		log.Println("Waiting for Postgres")
		time.Sleep(time.Second * 2)
	}

	return nil
}
