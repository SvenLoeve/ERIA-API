package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"viseh-api/database/ent"

	_ "github.com/lib/pq"
)

var Client *ent.Client

func DB() {
	err := *new(error)

	// Parse database port from string to int
	port, parseError := strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 32)
	if parseError != nil {
		log.Fatalf("failed parsing .env -> port: %v", parseError)
	}
	var host = os.Getenv("POSTGRES_HOST")
	var user = os.Getenv("POSTGRES_USER")
	var dbname = os.Getenv("POSTGRES_DB")
	var password = os.Getenv("POSTGRES_PASSWORD")
	var ssl = os.Getenv("SSL_MODE")
	var sslRootCert = os.Getenv("SSL_ROOT_CERT")
	var sslKey = os.Getenv("SSL_KEY")
	var sslCert = os.Getenv("SSL_CERT")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s sslrootcert=%s sslkey=%s sslcert=%s",
		host, port, user, password, dbname, ssl, sslRootCert, sslKey, sslCert)

	// Open database connection
	Client, err = ent.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Successfully connected to Postgres database!")
}
