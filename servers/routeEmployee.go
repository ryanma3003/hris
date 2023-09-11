package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
)

func InitializeRoutesEmployee(g *gin.RouterGroup) {
	{
		// employee
		g.GET("/emps", controllers.EmployeeIndex)
		g.GET("/emp/:id", controllers.EmployeeShow)

		// grade
		g.GET("/grades", controllers.GradeIndex)
		g.POST("/grade", controllers.GradeCreate)
		g.GET("/grade/:id", controllers.GradeShow)
		g.PUT("/grade/:id", controllers.GradeUpdate)
		g.DELETE("/grade/:id", controllers.GradeDelete)
	}
}
