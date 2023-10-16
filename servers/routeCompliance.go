package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
	"github.com/ryanma3003/hris/middleware"
)

func InitializeRoutesCompliance(g *gin.RouterGroup) {
	g.GET("/GenerateSlip", middleware.Authorize("resource", "*"), controllers.GenerateSlip)
	g.GET("/ListSlip", middleware.Authorize("resource", "*"), controllers.SalarySlipShow)
	g.GET("/ListSlipDetail/:id", middleware.Authorize("resource", "*"), controllers.SalarySlipDetailShow)
	g.POST("/CreateSlip/", middleware.Authorize("resource", "*"), controllers.SalarySlipCreate)
	g.POST("/CreateSlipDetail/", middleware.Authorize("resource", "*"), controllers.SalarySlipDetailCreate)
	g.PUT("/UpdateSlip/:id", middleware.Authorize("resource", "*"), controllers.SalarySlipUpdate)
	g.PUT("/UpdateSlipDetail/:id", middleware.Authorize("resource", "*"), controllers.SalarySlipDetailUpdate)
	g.PUT("/ApproveHR/:id", middleware.Authorize("resource", "*"), controllers.ApproveSlipHR)
	g.PUT("/ApproveFinance/:id", middleware.Authorize("resource", "*"), controllers.ApproveSlipFinance)
	g.DELETE("/DeleteSlip/:id", middleware.Authorize("resource", "*"), controllers.SalarySlipDelete)
	g.DELETE("/DeleteSlipDetail/:id", middleware.Authorize("resource", "*"), controllers.SalarySlipDetailDelete)
}
