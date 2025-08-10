package app

import (
	"demo/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *database.DB
	Router *gin.Engine
}

func NewApp(db *database.DB) *App {
	return &App{
		DB:     db,
		Router: gin.Default(),
	}
}

func (a *App) MustStart() {
	a.Router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (app *App) Run() error {

	if err := app.SetupRoutes(); err != nil {
		log.Fatal("Failed to setup server routes", "error", err)
		return err
	}

	if err := app.Router.Run(); err != nil {
		log.Fatal("Failed to start server", "error", err)
		return err
	}
	return nil
}

func (app *App) SetupRoutes() error {
	app.Router.GET("/order/:id")
	return nil
}

func (app *App) Stop() {
	log.Fatal("stoping backend service")
}
