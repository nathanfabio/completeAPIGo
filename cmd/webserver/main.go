package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanfabio/completeAPIGo/config/env"
	"github.com/nathanfabio/completeAPIGo/config/logger"
	"github.com/nathanfabio/completeAPIGo/internal/database"
	"github.com/nathanfabio/completeAPIGo/internal/database/sqlc"
	"github.com/nathanfabio/completeAPIGo/internal/handler/routes"
	"github.com/nathanfabio/completeAPIGo/internal/handler/userhandler"
	"github.com/nathanfabio/completeAPIGo/internal/repository/userepository"
	"github.com/nathanfabio/completeAPIGo/internal/service/userservice"
)

func main() {
	logger.InitLogger()
	slog.Info("Starting API")

	_, err := env.LoadingConfig(".")
	if err != nil {
		slog.Error("failedto load enviroment variables", err, slog.String("package", "main"))
		return
	}

	dbConnection, err := database.NewDBConnection()
	if err != nil {
		slog.Error("error to connect to database", "err", err, slog.String("package", "main"))
		return
	}

	router := chi.NewRouter()
	queries := sqlc.New(dbConnection)


	userRepo := userepository.NewUserRepository(dbConnection, queries)
	newUserService := userservice.NewUserService(userRepo)
	newUserHandler := userhandler.NewUserHandler(newUserService)


	routes.InitUserRoutes(router, newUserHandler)

	port := fmt.Sprintf(":%s", env.Env.GoPort)
	slog.Info(fmt.Sprintf("server running on port %s", port))
	err = http.ListenAndServe(port, router)
	if err != nil {
		slog.Error("error to start server", err, slog.String("package", "main"))
	}


}

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}


func (u user) LogUser() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Int("age", u.Age),
		slog.String("password", "*****"),
	)
}