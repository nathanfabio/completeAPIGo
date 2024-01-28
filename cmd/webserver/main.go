package main

import (
	"log/slog"

	"github.com/nathanfabio/completeAPIGo/config/logger"
)

func main() {
	logger.InitLogger()

	user := user{
		Name:     "nathan",
		Age:      24,
		Password: "1234",
	}
	
	slog.Info("Starting API")
	slog.Info("Creating a user", "user", user.LogUser())
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