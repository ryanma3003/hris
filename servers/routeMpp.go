package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
)

func InitializeRoutesMpp(g *gin.RouterGroup) {
	// MPP
	g.GET("/index/:employeeid/", controllers.MppIndex)
	g.POST("/createMpp", controllers.MppCreate)
	g.PUT("/updateMpp/:id", controllers.MppUpdate)
	g.GET("/showMpp/:id", controllers.MppShow)

	// REQ HEADCOUNT
	g.GET("/index/:employeeid/:period", controllers.ListMpp)
	g.GET("/formHeadcount/:mppid", controllers.FormHeadcount) // Form Create
	g.POST("/createHeadcount", controllers.CreateHeadcount)
	g.GET("/showAllHeadcount/:employeeid", controllers.ShowAllHeadcount)
	g.PUT("/updateReqHeadcount/:id", controllers.UpdateHeadcount)

	// HR
	g.GET("/listUnapproveMpp", controllers.MppListUnapprove)
	g.PUT("/approveMpp/:id", controllers.ApproveMpp)
	g.PUT("/revisionMpp/:id", controllers.RevisionMpp)
	g.PUT("/approveReqHeadcount/:id", controllers.ApproveReqHeadcount)
	g.PUT("/revisionReqHeadcount/:id", controllers.RevisionReqHeadcount)
}
