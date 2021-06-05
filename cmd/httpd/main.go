package main

import (
	"wmi-item-service/internal/config"
	"wmi-item-service/internal/httpd"
	"wmi-item-service/internal/repository"
	"wmi-item-service/internal/core/service"
	"wmi-item-service/pkg/postgresql"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/database/postgres"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.LoadConfig()

	router := gin.Default()

	dbConn, err := postgresql.NewConnection(
		cfg.Database.Host,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.Port,
		cfg.Database.Schema,
	)

	userRepo := repository.NewUserRepo(dbConn)
	userService := service.NewUserService(userRepo)

	residenceRepo := repository.NewResidenceRepo(dbConn)
	residenceService := service.NewResidenceService(residenceRepo)

	server := httpd.NewServer(router, userService, residenceService)

	err = server.Run()
	if err != nil {
		return err
	}

	return nil
}
