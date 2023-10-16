package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
	"github.com/ryanma3003/hris/middleware"
)

func InitializeRoutesMpp(g *gin.RouterGroup) {
	// MPP
	g.GET("/listMpp", middleware.Authorize("resource", "*"), controllers.MppIndex)
	g.POST("/createMpp", middleware.Authorize("resource", "*"), controllers.MppCreate)
	g.PUT("/updateMpp/:id", middleware.Authorize("resource", "*"), controllers.MppUpdate)
	g.GET("/showMpp/:id", middleware.Authorize("resource", "*"), controllers.MppShow)

	// REQ HEADCOUNT
	g.GET("/listHeadcount", middleware.Authorize("resource", "*"), controllers.ListMpp)
	g.GET("/formHeadcount/:mppid", middleware.Authorize("resource", "*"), controllers.FormHeadcount) // Form Create
	g.POST("/createHeadcount", middleware.Authorize("resource", "*"), controllers.CreateHeadcount)
	g.GET("/showAllHeadcount/:employeeid", middleware.Authorize("resource", "*"), controllers.ShowAllHeadcount)
	g.PUT("/updateReqHeadcount/:id", middleware.Authorize("resource", "*"), controllers.UpdateHeadcount)

	// HR
	g.GET("/listUnapproveMpp", middleware.Authorize("resource", "*"), controllers.MppListUnapprove)
	g.PUT("/approveMpp/:id", middleware.Authorize("resource", "*"), controllers.ApproveMpp)
	g.PUT("/revisionMpp/:id", middleware.Authorize("resource", "*"), controllers.RevisionMpp)
	g.PUT("/approveReqHeadcount/:id", middleware.Authorize("resource", "*"), controllers.ApproveReqHeadcount)
	g.PUT("/revisionReqHeadcount/:id", middleware.Authorize("resource", "*"), controllers.RevisionReqHeadcount)
}
