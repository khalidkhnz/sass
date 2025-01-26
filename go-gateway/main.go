package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/khalidkhnz/sass/echo-gateway/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitEnv()

	e := echo.New()

	fmt.Println("API GATEWAY IS UP AND RUNNING")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Environment Variables
	blogService := os.Getenv("BLOG_SERVICE_URL")
	ecomService := os.Getenv("ECOM_SERVICE_URL")
	sassService := os.Getenv("SASS_SERVICE_URL")

	if blogService == "" || ecomService == "" || sassService == "" {
		log.Fatal("Service URLs must be set in environment variables")
	}

	// Setup Proxy Groups
	setupProxy(e, "/go-blog", blogService)
	setupProxy(e, "/go-ecom", ecomService)
	setupProxy(e, "/go-sass", sassService)

	// Health Check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	// Graceful Shutdown
	go func() {
		if err := e.Start(config.GetPort()); err != nil && err != http.ErrServerClosed {
			log.Fatal("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Failed to gracefully shut down the server:", err)
	}
}

func setupProxy(e *echo.Echo, groupPath, serviceURL string) {
	target, err := url.Parse(serviceURL)
	if err != nil {
		e.Logger.Fatal(err)
	}
	group := e.Group(groupPath)
	group.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{URL: target},
		}),
		Rewrite: map[string]string{
			groupPath:      "/",
			groupPath + "/*": "/$1",
		},
	}))
}
