package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"saas-task-manager/config"
	"saas-task-manager/models"
	"saas-task-manager/utils"
)

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	hashed, _ := utils.HashPassword(user.Password)
	user.Password = hashed
	user.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := config.UserCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created", "id": result.InsertedID})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := config.UserCollection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := utils.GenerateToken(user.ID.Hex())
	c.JSON(http.StatusOK, gin.H{"token": token})
}
