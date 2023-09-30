package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
	"github.com/ryanma3003/hris/middleware"
)

func InitializeRoutesEmployee(g *gin.RouterGroup) {
	{
		// employee
		g.GET("/emps", controllers.EmployeeIndex)
		g.POST("/emp", controllers.EmployeeCreate)
		g.GET("/emp/:id", controllers.EmployeeShow)
		g.PUT("/emp/:id", controllers.EmployeeUpdate)
		g.DELETE("/emp/:id", middleware.Authorize("resource", "*"), controllers.EmployeeDelete)

		// role
		g.GET("/roles", controllers.RoleIndex)

		// grade
		g.GET("/grades", controllers.GradeIndex)
		g.POST("/grade", middleware.Authorize("resource", "*"), controllers.GradeCreate)
		g.GET("/grade/:id", controllers.GradeShow)
		g.PUT("/grade/:id", middleware.Authorize("resource", "*"), controllers.GradeUpdate)
		g.DELETE("/grade/:id", middleware.Authorize("resource", "*"), controllers.GradeDelete)

		// jobdesc
		g.GET("/jobdescs", controllers.JobDescriptionIndex)
		g.POST("/jobdesc", middleware.Authorize("resource", "*"), controllers.JobDescriptionCreate)
		g.GET("/jobdesc/:id", controllers.JobDescriptionShow)
		g.PUT("/jobdesc/:id", middleware.Authorize("resource", "*"), controllers.JobDescriptionUpdate)
		g.DELETE("/jobdesc/:id", middleware.Authorize("resource", "*"), controllers.JobDescriptionDelete)

		// level
		g.GET("/levels", controllers.LevelIndex)
		g.POST("/level", middleware.Authorize("resource", "*"), controllers.LevelCreate)
		g.GET("/level/:id", controllers.LevelShow)
		g.PUT("/level/:id", middleware.Authorize("resource", "*"), controllers.LevelUpdate)
		g.DELETE("/level/:id", middleware.Authorize("resource", "*"), controllers.LevelDelete)

		// sup
		g.GET("/sups", controllers.SupervisionIndex)
		g.POST("/sup", middleware.Authorize("resource", "*"), controllers.SupervisionCreate)
		g.GET("/sup/:id", controllers.SupervisionShow)
		g.PUT("/sup/:id", middleware.Authorize("resource", "*"), controllers.SupervisionUpdate)
		g.DELETE("/sup/:id", middleware.Authorize("resource", "*"), controllers.SupervisionDelete)

		// dep
		g.GET("/deps", controllers.DepartmentIndex)
		g.POST("/dep", middleware.Authorize("resource", "*"), controllers.DepartmentCreate)
		g.GET("/dep/:id", controllers.DepartmentShow)
		g.PUT("/dep/:id", middleware.Authorize("resource", "*"), controllers.DepartmentUpdate)
		g.DELETE("/dep/:id", middleware.Authorize("resource", "*"), controllers.DepartmentDelete)

		// div
		g.GET("/divs", controllers.DivisionIndex)
		g.POST("/div", middleware.Authorize("resource", "*"), controllers.DivisionCreate)
		g.GET("/div/:id", controllers.DivisionShow)
		g.PUT("/div/:id", middleware.Authorize("resource", "*"), controllers.DivisionUpdate)
		g.DELETE("/div/:id", middleware.Authorize("resource", "*"), controllers.DivisionDelete)

		// ptkp
		g.GET("/ptkps", controllers.PtkpIndex)
		g.POST("/ptkp", middleware.Authorize("resource", "*"), controllers.PtkpCreate)
		g.GET("/ptkp/:id", controllers.PtkpShow)
		g.PUT("/ptkp/:id", middleware.Authorize("resource", "*"), controllers.PtkpUpdate)
		g.DELETE("/ptkp/:id", middleware.Authorize("resource", "*"), controllers.PtkpDelete)

		// pph
		g.GET("/pphs", controllers.PphIndex)
		g.POST("/pph", middleware.Authorize("resource", "*"), controllers.PphCreate)
		g.GET("/pph/:id", controllers.PphShow)
		g.PUT("/pph/:id", middleware.Authorize("resource", "*"), controllers.PphUpdate)
		g.DELETE("/pph/:id", middleware.Authorize("resource", "*"), controllers.PphDelete)
	}
}
