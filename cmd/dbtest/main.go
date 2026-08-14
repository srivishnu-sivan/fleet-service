package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Changed localhost to 127.0.0.1 to avoid Windows IPv6 loopback errors
	dsn := "postgres://fleet_user:fleet_password@127.0.0.1:5439/fleet?sslmode=disable"
		
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	var result int
	err = pool.QueryRow(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB CONNECTED:", result)
}
