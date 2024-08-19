package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnection = ""
	Port             = 0
)

func Loading() {
	var err error
	if err = godotenv.Load(); err != nil {
		slog.Error("| NOT LOAD .ENV | ERROR:", err)
		return
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		slog.Warn("| NOT LOAD PORT | NEW PORT DEFAULT :9000 | ERROR:", err)
		Port = 9000
	}

	StringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
}
