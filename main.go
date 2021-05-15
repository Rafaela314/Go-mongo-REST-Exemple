package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/config"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/logs"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/middleware"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
	m "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func welcome() {
	//https://patorjk.com/software/taag/#p=display&f=Slant&t=golang%20microservice

	fmt.Println("	__                               _                                      _         ")
	fmt.Println("	____ _____  / /___ _____  ____ _   ____ ___  (_)_____________  ________  ______   __(_)_______ ")
	fmt.Println("   / __ `/ __ \\/ / __ `/ __ \\/ __ `/  / __ `__ \\/ / ___/ ___/ __ \\/ ___/ _ \\/ ___/ | / / / ___/ _ \\")
	fmt.Println("  / /_/ / /_/ / / /_/ / / / / /_/ /  / / / / / / / /__/ /  / /_/ (__  )  __/ /   | |/ / / /__/  __/")
	fmt.Println("  \\__, /\\____/_/\\__,_/_/ /_/\\__, /  /_/ /_/ /_/_/\\___/_/   \\____/____/\\___/_/    |___/_/\\___/\\___/ ")
	fmt.Println(" /____/                    /____/                                                                  ")
}

// @title Star Wars API
// @Version 1.0
// @BasePath /

func main() {

	welcome()

	//load configuration
	cfg := config.LoadConfig()

	fmt.Printf("/n CFG %v /n", cfg)

	//server configuration
	e := echo.New()
	e.HideBanner = true
	e.Logger = logs.LogrusAdapter()
	e.Logger.SetLevel(cfg.LogLevel)
	e.Logger.SetOutput(cfg.LogOutput)

	//Middlerware
	e.Use(middleware.Logger())

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// CORS config
	allowOrigins := []string{"*"}
	if s, ok := os.LookupEnv("SERVER_CORS"); ok {
		allowOrigins = []string{s}
	}

	e.Use(m.CORSWithConfig(m.CORSConfig{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	//system test
	e.GET("/system/alive", isAlive)
	e.GET("/system/echo", getEcho)
	e.GET("/system/now", now)

	e.Logger.Fatal(e.Start(":1223"))
}

func isAlive(c echo.Context) error {
	return c.String(http.StatusOK, "Yes.")
}

func getEcho(c echo.Context) error {
	p := c.QueryParam("text")
	return c.String(http.StatusOK, p)
}

func now(c echo.Context) error {
	t := time.Now()
	ts := models.Timestamp{
		Unix:  t.UnixNano(),
		UTC:   t.UTC(),
		Local: t.Local(),
	}
	return c.JSON(http.StatusOK, ts)

}
