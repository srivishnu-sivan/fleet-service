package app

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func New() *App {
	router := gin.Default()
	return &App{
		Router: router,
	}
}
