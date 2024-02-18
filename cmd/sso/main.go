package main

import (
	"fmt"
	"log/slog"

	"sso/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}

func setupLogger(env string) *slog.Logger {

}
