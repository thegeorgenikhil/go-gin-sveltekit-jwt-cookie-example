package app

import (
	"context"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/internal/api/routes/auth"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/internal/api/routes/user"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/db"
)

type App struct {
	db *db.UserDB
}

func New(db *db.UserDB) *App {
	return &App{
		db: db,
	}
}

func (a *App) Run(ctx context.Context) error {
	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		},
	))

	auth.RegisterAuthRoutes(r, a.db)
	user.RegisterUserRoutes(r, a.db)

	return r.Run(fmt.Sprintf(":%d", 8080))
}
