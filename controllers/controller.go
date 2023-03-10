package controllers

import (
	"apirest-go-gin/database"
	"apirest-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayAllStudents(c *gin.Context) {
	var students []models.Stundents
	database.DB.Find(&students)
	c.JSON(200, students)
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hello " + name + "!",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Stundents
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}
