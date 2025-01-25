package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/khalidkhnz/sass/go-ecom/config"
	"github.com/khalidkhnz/sass/go-ecom/schemas"
	"github.com/khalidkhnz/sass/go-ecom/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



type APIServer struct {
	listenAddr string
	echo       *echo.Echo
	db         *pgx.Conn
	migrate    bool
}

func (s *APIServer) listenAndServe() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "Hello From Go-Ecom",
		})
	})

	if err := s.initDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	s.initMiddlewares()
	s.initRoutes()

	log.Printf("Server starting on %s", s.listenAddr)
	if err := s.echo.Start(s.listenAddr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}



func (s *APIServer) initMiddlewares() {
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.CORS())
}

func (s *APIServer) initRoutes() {
	// Add route groups and handlers here
	api := s.echo.Group("/api")
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})
}

func (s *APIServer) runMigration() {
	if !s.migrate {
		return
	}

	log.Println("Running database migrations...")
	
	// Example migration queries
	migrations := schemas.MigrationQueries

	for _, query := range migrations {
		_, err := s.db.Exec(context.TODO(), query)
		if err != nil {
			log.Printf("Migration failed: %v", err)
			continue
		}
	}

	log.Println("Database migrations completed")
}


func (s *APIServer) initDatabase() error {
	conn, err := services.ConnectToDb()
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	s.db = conn
	go s.runMigration()
	return nil
}

func main() {
	config.InitEnv()
	echo := echo.New()

	server := &APIServer{
		listenAddr: config.GetPort(),
		echo:       echo,
		migrate: false,
	}

	server.listenAndServe()
}