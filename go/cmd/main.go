package main

import (
	"main.go/internal/app"
	"main.go/internal/http"
	"main.go/internal/store"
)

func main() {
	store := store.New()

	app := app.New(store)

	api := http.New(app)
	api.Engine()
}
