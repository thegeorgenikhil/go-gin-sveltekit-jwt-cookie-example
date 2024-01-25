package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/db"
)

type userController struct {
	db *db.UserDB
}

func newUserController(db *db.UserDB) *userController {
	return &userController{
		db: db,
	}
}

func (uc *userController) getMyId(c *gin.Context) {
	email := c.GetString("email")
	user, _ := uc.db.GetUser(email)

	c.JSON(200, gin.H{
		"id": user.ID,
	})
}
