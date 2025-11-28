package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"zotes/servers/web/internal/app"
	"zotes/shared/create"
	"zotes/shared/structs"

	"github.com/joho/godotenv"
)

func main() {
	logger := create.CreateLogger()

	var err error
	if os.Getenv("ENVIRONMENT") != "prod" {
		err = godotenv.Load("../../envs/web.env")
		if err != nil {
			logger.Error("Error loading .env file", slog.Any("error", err))
		}
		err = godotenv.Load("../../envs/environment.env")
		if err != nil {
			logger.Error("Error loading .env file", slog.Any("error", err))
		}
		err = godotenv.Load("../../envs/s3.env")
		if err != nil {
			logger.Error("Error loading .env file", slog.Any("error", err))
		}
		err = godotenv.Load("../../envs/pdfworker.env")
		if err != nil {
			logger.Error("Error loading .env file", slog.Any("error", err))
		}
	}

	if os.Getenv("ENVIRONMENT") == "dev" {
		logger.Debug("Reminder to start docker services with 'make on' at root, if database hasn't been setup use 'make up' to migrate.")
	}

	addr := flag.String(("addr"), ":4444", "HTTP network address")

	dbauth := &structs.DBAuth{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	pool := create.PGPool(dbauth)
	app := &app.App{
		DB:                pool,
		Logger:            logger,
		Session:           create.PGSessionManager(pool, "zotes-client"),
		OidcAuthenticator: create.OidcAuthenticator("CLIENT_UI"),
		// JwtIssuer:         create.JwtIssuer(""),
		// JwtValidator:      create.JwtValidator(""),
		Rbac: create.Casbin(dbauth),
	}

	defer pool.Close()
	routes := GetRoutes(app)
	// InsertRoutes(routes, deps)
	err = http.ListenAndServe(*addr, routes)

	app.Logger.Error("Server stopped unexpectedly" + err.Error())

}
