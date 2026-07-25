package main

import (
	"context"
	"dbApi/internal/app"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	_ "dbApi/docs"
)

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
