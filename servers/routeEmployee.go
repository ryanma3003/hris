package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/middleware"
)

func InitializeRoutesEmployee(g *gin.RouterGroup) {
	{
		// employee
		g.GET("/emps", controllers.EmployeeIndex)
		g.POST("/emp", controllers.EmployeeCreate)
		g.GET("/emp/:id", controllers.EmployeeShow)
		g.PUT("/emp/:id", controllers.EmployeeUpdate)
		g.DELETE("/emp/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EmployeeDelete)

		// grade
		g.GET("/grades", middleware.Authorize("resource", "*", db.Adapter), controllers.GradeIndex)
		g.POST("/grade", middleware.Authorize("resource", "*", db.Adapter), controllers.GradeCreate)
		g.GET("/grade/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.GradeShow)
		g.PUT("/grade/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.GradeUpdate)
		g.DELETE("/grade/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.GradeDelete)

		// jobdesc
		g.GET("/jobdescs", middleware.Authorize("resource", "*", db.Adapter), controllers.JobDescriptionIndex)
		g.POST("/jobdesc", middleware.Authorize("resource", "*", db.Adapter), controllers.JobDescriptionCreate)
		g.GET("/jobdesc/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.JobDescriptionShow)
		g.PUT("/jobdesc/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.JobDescriptionUpdate)
		g.DELETE("/jobdesc/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.JobDescriptionDelete)

		// level
		g.GET("/levels", middleware.Authorize("resource", "*", db.Adapter), controllers.LevelIndex)
		g.POST("/level", middleware.Authorize("resource", "*", db.Adapter), controllers.LevelCreate)
		g.GET("/level/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.LevelShow)
		g.PUT("/level/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.LevelUpdate)
		g.DELETE("/level/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.LevelDelete)

		// sup
		g.GET("/sups", middleware.Authorize("resource", "*", db.Adapter), controllers.SupervisionIndex)
		g.POST("/sup", middleware.Authorize("resource", "*", db.Adapter), controllers.SupervisionCreate)
		g.GET("/sup/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.SupervisionShow)
		g.PUT("/sup/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.SupervisionUpdate)
		g.DELETE("/sup/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.SupervisionDelete)

		// dep
		g.GET("/deps", middleware.Authorize("resource", "*", db.Adapter), controllers.DepartmentIndex)
		g.POST("/dep", middleware.Authorize("resource", "*", db.Adapter), controllers.DepartmentCreate)
		g.GET("/dep/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.DepartmentShow)
		g.PUT("/dep/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.DepartmentUpdate)
		g.DELETE("/dep/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.DepartmentDelete)

		// div
		g.GET("/divs", middleware.Authorize("resource", "*", db.Adapter), controllers.DivisionIndex)
		g.POST("/div", middleware.Authorize("resource", "*", db.Adapter), controllers.DivisionCreate)
		g.GET("/div/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.DivisionShow)
		g.PUT("/div/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.DivisionUpdate)
		g.DELETE("/div/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.DepartmentDelete)
	}
}
