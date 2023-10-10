package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
)

func InitializeRoutesService(g *gin.RouterGroup) {
	g.GET("/paidleaves", controllers.PaidleaveIndex)
	g.GET("/paidleave/:id", controllers.PaidleaveShow)
	g.POST("/paidleave/create", controllers.PaidleaveCreate)
	g.PUT("/paidleave/update/:id", controllers.PaidleaveUpdate)
	g.DELETE("/paidleave/delete/:id", controllers.PaidleaveDelete)

	// Tambahkan endpoint baru untuk melihat jumlah data cuti dengan Status=1 untuk setiap NikID
	g.GET("/count-status/:id", controllers.PaidleaveCountStatus)

	// Loan
	g.GET("/loans", controllers.LoanIndex)
	g.GET("/loan/:id", controllers.LoanShow)
	g.POST("/loan/create", controllers.LoanCreate)
	g.PUT("/loan/update/:id", controllers.LoanUpdate)
	g.DELETE("/loan/delete/:id", controllers.LoanDelete)

	// Insurance
	g.GET("/insurances", controllers.InsuranceIndex)
	g.GET("/insurance/:id", controllers.InsuranceShow)
	g.POST("/insurance/create", controllers.InsuranceCreate)
	g.PUT("/insurance/update/:id", controllers.InsuranceUpdate)
	g.DELETE("/insurance/delete/:id", controllers.InsuranceDelete)

	// Asset
	g.GET("/assets", controllers.AssetIndex)
	g.GET("/asset/:id", controllers.AssetShow)
	g.POST("/asset/create", controllers.CreateAsset)
	g.PUT("/asset/update/:id", controllers.AssetUpdate)
	g.DELETE("/asset/delete/:id", controllers.AssetDelete)
}
