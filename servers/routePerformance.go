package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
	"github.com/ryanma3003/hris/db"
	"github.com/ryanma3003/hris/middleware"
)

func InitializeRoutesPerformance(g *gin.RouterGroup) {
	// result
	g.GET("/perfs", controllers.EvaluationIndex)
	// g.POST("/emp", controllers.EmployeeCreate)
	// g.GET("/emp/:id", controllers.EmployeeShow)
	// g.PUT("/emp/:id", controllers.EmployeeUpdate)
	// g.DELETE("/emp/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EmployeeDelete)

	// evalform
	g.GET("/evalforms", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationFormIndex)
	g.POST("/evalform", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationFormCreate)
	g.GET("/evalform/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationFormShow)
	g.PUT("/evalform/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationFormUpdate)
	g.DELETE("/evalform/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationFormDelete)

	// evalpoint
	g.GET("/evalpoints", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationPointIndex)
	g.POST("/evalpoint", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationPointCreate)
	g.GET("/evalpoint/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationPointShow)
	g.PUT("/evalpoint/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationPointUpdate)
	g.DELETE("/evalpoint/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.EvaluationPointDelete)

	// selfeval
	g.GET("/selfevals", controllers.SelfPerformanceIndex)
	g.POST("/selfeval", controllers.SelfPerformanceCreate)
	g.GET("/selfeval/:id", controllers.SelfPerformanceShow)
	g.PUT("/selfeval/:id", controllers.SelfPerformanceUpdate)
	g.DELETE("/selfeval/:id", middleware.Authorize("resource", "*", db.Adapter), controllers.SelfPerformanceDelete)
}
