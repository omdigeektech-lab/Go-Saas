package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"saas-task-manager/config"
	"saas-task-manager/models"
)

func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")
	objID, _ := primitive.ObjectIDFromHex(userID.(string))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.TaskCollection.Find(ctx, bson.M{"user_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching tasks"})
		return
	}

	var tasks []models.Task
	cursor.All(ctx, &tasks)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	task.UserID, _ = primitive.ObjectIDFromHex(userID.(string))
	task.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task created"})
}

func UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
