package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/internal/api/middleware"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/db"
)

func RegisterUserRoutes(r *gin.Engine, db *db.UserDB) {
	controller := newUserController(db)
	routes := r.Group("/user")
	routes.Use(middleware.AuthMiddleware())

	routes.GET("/get-my-id", controller.getMyId)
}
