package routes

import (
	"apirest-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.DisplayAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.POST("/students", controllers.CreateStudent)
	r.Run(":5000")
}
