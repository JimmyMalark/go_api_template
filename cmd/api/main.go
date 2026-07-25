package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/jimmymalark/go_api_template/internal/app"

	_ "github.com/jimmymalark/go_api_template/docs"
)

// @title           Go API Template
// @version         1.0
// @description     A production-ready Go API template.
// @BasePath        /v1
func main() {
	a, err := app.New(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", a.Config.App.Host, a.Config.App.Port)

	slog.Info("server started",
		"address", addr,
	)

	log.Fatal(http.ListenAndServe(addr, a.Router))
}
