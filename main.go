package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	SQL_PASSWORD := os.Getenv("SQL_PASSWORD")

	dsn := fmt.Sprintf("postgresql://celia:%v@free-tier14.aws-us-east-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster%%3Dsnow-kudu-6708", SQL_PASSWORD)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	defer conn.Close(context.Background())

	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
