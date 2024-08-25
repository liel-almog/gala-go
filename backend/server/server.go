package server

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/liel-almog/gala-go/backend/configs"
)

var addr = ":8080"

var app *echo.Echo

func Serve() {
	app = echo.New()

	app.Use(middleware.Recover())

	setupRouter(app)

	port, err := configs.GetEnv("PORT")
	if err != nil {
		log.Fatal("GetEnv: ", err)
	}

	if port != "" {
		addr = fmt.Sprintf(":%s", port)
	}

	fmt.Println("Server strating on port", addr)

	if err := app.Start(addr); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Shutdown(ctx context.Context) error {
	err := app.Shutdown(ctx)

	if err != nil {
		return err
	}

	return nil
}
