package main

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	app "github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/internal/api/app"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/db"
)

const file = "user.db"

func main() {
	userDb, err := db.New(file)
	if err != nil {
		log.Fatalln("db new failed: ", err)
	}

	a := app.New(userDb)

	if err := a.Run(context.Background()); err != nil {
		log.Fatalln("app run failed: ", err)
	}
}
