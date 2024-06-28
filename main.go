package main

import (
	"dependency-injection/domain/user"
	"dependency-injection/repository/inmemory"
	"dependency-injection/usecase/register"
	"log/slog"
	"os"
)

func main() {
	inmemoryRepository := inmemory.NewRepository()
	registerService := register.NewService(inmemoryRepository)

	//       ↑
	// Switch injection
	//       ↓

	// mysqlRepository := mysql.NewRepository(
	// 	mysql.DataSource("user", "password", "127.0.0.1", 3306, "db"),
	// )
	// registerService := register.NewService(mysqlRepository)

	err := registerService.Register(user.User{Name: "hoge"})
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("success")
}
