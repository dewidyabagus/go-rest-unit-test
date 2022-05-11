package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"

	"go-api-v1/api"
	"go-api-v1/api/v1/user"
	"go-api-v1/api/v1/welcome"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error Loding Env:", err.Error())
	}

	e := echo.New()

	routes := &api.Routes{
		Welcome: welcome.NewController(),
		User:    user.NewController(),
	}
	api.RegisterRoutes(e, routes)

	go func() {
		listen := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
		if err := e.Start(listen); err != nil {
			log.Println("Shutting Down REST Server")
			os.Exit(0)
		}
	}()

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println("Error Shutting Down:", err.Error())
	}
}
