package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/db"
)

func RegisterAuthRoutes(r *gin.Engine, db *db.UserDB) {
	controller := newAuthController(db)
	routes := r.Group("/auth")

	routes.POST("/signup", controller.signup)
	routes.POST("/login", controller.login)
}
