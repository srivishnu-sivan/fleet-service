package main

import (
	"context"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/srivishnu-sivan/fleet-service/internal/app"
	internal "github.com/srivishnu-sivan/fleet-service/internal/config"
	"github.com/srivishnu-sivan/fleet-service/internal/database"
)

func main() {
	cfg := internal.Load()

	ctx := context.Background()

connString :=  "postgres://fleet_user:fleet_password@127.0.0.1:5439/fleet?sslmode=disable"

pool, err := database.NewPostgresPool(ctx, connString)
if err != nil {
    log.Fatal(err)
}
defer pool.Close()

application := app.New(pool)

if err := application.Router.Run(":" + cfg.Port); err != nil {
    log.Fatal(err)
}

}
