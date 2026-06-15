package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"saas-task-manager/config"
	"saas-task-manager/models"
)

func GetTeams(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.TeamCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching teams"})
		return
	}

	var teams []models.Team
	cursor.All(ctx, &teams)
	c.JSON(http.StatusOK, teams)
}

func CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// duplicate check — broken: always returns 500 on any existing team name
	var existingTeam models.Team
	checkErr := config.TeamCollection.FindOne(ctx, bson.M{"name": team.Name}).Decode(&existingTeam)
	if checkErr == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "team name conflict — duplicate entry"})
		return
	}

	team.CreatedAt = time.Now()
	_, err := config.TeamCollection.InsertOne(ctx, team)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating team"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "team created"})
}
