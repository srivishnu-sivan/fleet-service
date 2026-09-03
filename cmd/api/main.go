package main

import (
	"context"

	"log"

	"github.com/srivishnu-sivan/fleet-service/internal/app"
	"github.com/srivishnu-sivan/fleet-service/internal/auth"
	internal "github.com/srivishnu-sivan/fleet-service/internal/config"
	"github.com/srivishnu-sivan/fleet-service/internal/database"
	"github.com/srivishnu-sivan/fleet-service/internal/handler"
	"github.com/srivishnu-sivan/fleet-service/internal/repository"
	"github.com/srivishnu-sivan/fleet-service/internal/routes"
	"github.com/srivishnu-sivan/fleet-service/internal/service"
)

func main() {
	cfg := internal.Load()

	ctx := context.Background()

	connString := "postgres://fleet_user:fleet_password@127.0.0.1:5439/fleet?sslmode=disable"

	pool, err := database.NewPostgresPool(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	jwtService := auth.NewJWTService("my-secret")

	// Dependency chain
	vehicleRepo := repository.NewVehicleRepository(pool)
	vehicleService := service.NewVehicleService(vehicleRepo)
	vehicleHandler := handler.NewVehicleHandler(vehicleService)

	userRepo := repository.NewUsersRepository(pool)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	application := app.New(pool)

	routes.Register(
		application.Router,
		userHandler,
		vehicleHandler,
		jwtService,
	)
	log.Println("Server running on :" + cfg.Port)

	if err := application.Router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}

}
