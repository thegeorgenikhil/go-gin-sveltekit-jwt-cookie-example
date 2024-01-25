package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/bcrypt"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/db"
	"github.com/thegeorgenikhil/go-gin-sveltekit-jwt-cookie-example/pkg/jwt"
)

type authController struct {
	db *db.UserDB
}

func newAuthController(db *db.UserDB) *authController {
	return &authController{
		db: db,
	}
}

type signupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *authController) signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	exists, err := ac.db.CheckIfUserExists(req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "unable to check if user exists"})
		return
	}

	if exists {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "unable to hash password"})
		return
	}

	if err := ac.db.CreateUser(req.Email, hashedPassword); err != nil {
		c.JSON(400, gin.H{"error": "unable to create user"})
		return
	}

	token, _ := jwt.GenerateToken(req.Email, 24*time.Hour)

	c.JSON(201, gin.H{
		"message": "user created",
		"token":   token,
	})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *authController) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	user, err := ac.db.GetUser(req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "user with this email does not exist"})
		return
	}

	if ok := bcrypt.VerifyPassword(user.Password, req.Password); !ok {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	token, _ := jwt.GenerateToken(req.Email, 24*time.Hour)

	c.JSON(200, gin.H{
		"message": "user logged in",
		"token":   token,
	})
}
