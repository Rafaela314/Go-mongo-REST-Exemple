package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/config"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/logs"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/settings"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/middleware"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/router"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"
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
	/*
		// database connection
		db, err := openDBConnection(cfg)
		if err != nil {
			os.Exit(1)
		}
		defer db.Close()
	*/

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

	//create server
	s := server.NewServer(e, settings.NewSettings(cfg.SwapiURL))

	//API routes
	r := router.New(s)
	r.Setup()

	//start server
	e.Logger.Fatal(s.Start(":" + strconv.FormatUint(cfg.Port, 10)))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := s.E.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Server Shutdown Failed", err.Error())
	}
	cancel()

}

/*
func openDBConnection(cfg config) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	for i := 0; i < 5; i++ {
		if i > 0 {
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		fmt.Printf("Connecting to database (tries=%d)... ", i+1)
		db, err = sqlx.Open(cfg.DBDriver, cfg.DBURL)
		if err != nil {
			fmt.Printf("ERROR!\n%v\n\n", err)
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Printf("ERROR!\ndatabase error: %v\n", err)
		} else {
			break
		}
	}

	return db, err
}*/
