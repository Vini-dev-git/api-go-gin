package routes

import (
	"apirest-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.DisplayAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.GET("/students/:id", controllers.SearchStudentById)
	r.GET("/students/cpf/:cpf", controllers.SearcStudentByCPF)
	r.POST("/students", controllers.CreateStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.DELETE("/student/:id", controllers.DeletStudent)
	r.Run(":5000")
}
