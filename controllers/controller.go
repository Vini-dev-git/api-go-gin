package controllers

import "github.com/gin-gonic/gin"

func DisplayAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Mando",
	})
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hello " + name + "!",
	})
}
