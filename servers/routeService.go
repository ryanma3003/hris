package servers

<<<<<<< HEAD
import "github.com/gin-gonic/gin"

func InitializeRoutesService(g *gin.RouterGroup) {
=======
import (
	"github.com/gin-gonic/gin"
	"github.com/ryanma3003/hris/controllers"
)

func InitializeRoutesService(g *gin.RouterGroup) {
	g.GET("/", controllers.PaidleaveIndex)
	g.GET("/:id", controllers.PaidleaveShow)
	g.POST("/create", controllers.PaidleaveCreate)
	g.PUT("/update/:id", controllers.PaidleaveUpdate)
	g.DELETE("/delete/:id", controllers.PaidleaveDelete)

	// Tambahkan endpoint baru untuk melihat jumlah data cuti dengan Status=1 untuk setiap NikID
	g.GET("/count-status/:nikid", controllers.PaidleaveCountStatus)

	// Loan
	g.GET("/", controllers.LoanIndex)
	g.GET("/:id", controllers.LoanShow)
	g.POST("/create", controllers.LoanCreate)
	g.PUT("/update/:id", controllers.LoanUpdate)
	g.DELETE("/delete/:id", controllers.LoanDelete)

	// Insurance
	g.GET("/", controllers.InsuranceIndex)
	g.GET("/:id", controllers.InsuranceShow)
	g.POST("/create", controllers.InsuranceCreate)
	g.PUT("/update/:id", controllers.InsuranceUpdate)
	g.DELETE("/delete/:id", controllers.InsuranceDelete)

	// Asset
	g.GET("/", controllers.AssetIndex)
	g.GET("/:id", controllers.AssetShow)
	g.POST("/create", controllers.CreateAsset)
	g.PUT("/update/:id", controllers.AssetUpdate)
	g.DELETE("/delete/:id", controllers.AssetDelete)
>>>>>>> f13f7b863e795beef5a9a954bafb384fb419b07d
}
