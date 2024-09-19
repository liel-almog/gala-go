package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/liel-almog/gala-go/backend/configs"
)

var addr = ":8080"

var server *http.Server

func Serve() {
	app := echo.New()

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

	server = &http.Server{
		Addr:              addr,
		Handler:           app,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Shutdown(ctx context.Context) error {
	err := server.Shutdown(ctx)

	if err != nil {
		return err
	}

	return nil
}
