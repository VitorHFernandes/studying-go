package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	config.Loading()
	slog.Info("| RUNNING SERVER ON PORT:", config.Port)

	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
