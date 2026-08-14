package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Router *gin.Engine
	DB *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *App {
	router := gin.Default()
	return &App{
		Router: router,
		DB: pool,
	}
}


