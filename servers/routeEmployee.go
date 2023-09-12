package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
)

func InitializeRoutesEmployee(g *gin.RouterGroup) {
	{
		// employee
		g.GET("/emps", controllers.EmployeeIndex)
		g.POST("/emp", controllers.EmployeeCreate)
		g.GET("/emp/:id", controllers.EmployeeShow)

		// grade
		g.GET("/grades", controllers.GradeIndex)
		g.POST("/grade", controllers.GradeCreate)
		g.GET("/grade/:id", controllers.GradeShow)
		g.PUT("/grade/:id", controllers.GradeUpdate)
		g.DELETE("/grade/:id", controllers.GradeDelete)

		// jobdesc
		g.GET("/jobdescs", controllers.JobDescriptionIndex)
		g.POST("/jobdesc", controllers.JobDescriptionCreate)
		g.GET("/jobdesc/:id", controllers.JobDescriptionShow)
		g.PUT("/jobdesc/:id", controllers.JobDescriptionUpdate)
		g.DELETE("/jobdesc/:id", controllers.JobDescriptionDelete)

		// level
		g.GET("/levels", controllers.LevelIndex)
		g.POST("/level", controllers.LevelCreate)
		g.GET("/level/:id", controllers.LevelShow)
		g.PUT("/level/:id", controllers.LevelUpdate)
		g.DELETE("/level/:id", controllers.LevelDelete)

		// sup
		g.GET("/sups", controllers.SupervisionIndex)
		g.POST("/sup", controllers.SupervisionCreate)
		g.GET("/sup/:id", controllers.SupervisionShow)
		g.PUT("/sup/:id", controllers.SupervisionUpdate)
		g.DELETE("/sup/:id", controllers.SupervisionDelete)

		// dep
		g.GET("/deps", controllers.DepartmentIndex)
		g.POST("/dep", controllers.DepartmentCreate)
		g.GET("/dep/:id", controllers.DepartmentShow)
		g.PUT("/dep/:id", controllers.DepartmentUpdate)
		g.DELETE("/dep/:id", controllers.DepartmentDelete)

		// div
		g.GET("/divs", controllers.DivisionIndex)
		g.POST("/div", controllers.DivisionCreate)
		g.GET("/div/:id", controllers.DivisionShow)
		g.PUT("/div/:id", controllers.DivisionUpdate)
		g.DELETE("/div/:id", controllers.DepartmentDelete)
	}
}
