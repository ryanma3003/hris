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

		// candidate
		g.GET("/cans", controllers.CandidateIndex)
		g.POST("/can", controllers.CandidateCreate)
		g.GET("/can/:id", controllers.CandidateShow)
		g.PUT("/can/:id", controllers.CandidateUpdate)
		g.DELETE("/can/:id", middleware.Authorize("resource", "*"), controllers.CandidateDelete)

		// role
		g.GET("/roles", controllers.RoleIndex)

		// upload avatar
		g.PUT("/avatar/:id", controllers.UpdateAvatar)
		g.GET("/avatar/:id", controllers.GetAvatar)

		// upload avatar candidate
		g.PUT("/avatar-candidate/:id", controllers.UpdateAvatarCandidate)
		g.GET("/avatar-candidate/:id", controllers.GetAvatarCandidate)

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

		// family
		g.GET("/familys", controllers.FamilyIndex)
		g.POST("/family", controllers.FamilyCreate)
		g.GET("/family/:id", controllers.FamilyShow)
		g.PUT("/family/:id", controllers.FamilyEdit)
		g.DELETE("/family/:id", controllers.FamilyDelete)

		// education
		g.GET("/educations", controllers.EducationIndex)
		g.POST("/education", controllers.EducationCreate)
		g.GET("/education/:id", controllers.EducationShow)
		g.PUT("/education/:id", controllers.EducationEdit)
		g.DELETE("/education/:id", controllers.EducationDelete)

		// course
		g.GET("/courses", controllers.CourseIndex)
		g.POST("/course", controllers.CourseCreate)
		g.GET("/course/:id", controllers.CourseShow)
		g.PUT("/course/:id", controllers.CourseEdit)
		g.DELETE("/course/:id", controllers.CourseDelete)

		// health
		g.GET("/healths", controllers.HealthDiseaseIndex)
		g.POST("/health", controllers.HealthDiseaseCreate)
		g.GET("/health/:id", controllers.HealthDiseaseShow)
		g.PUT("/health/:id", controllers.HealthDiseaseEdit)
		g.DELETE("/health/:id", controllers.HealthDiseaseDelete)

		// criminal
		g.GET("/criminals", controllers.CriminalNoteIndex)
		g.POST("/criminal", controllers.CriminalNoteCreate)
		g.GET("/criminal/:id", controllers.CriminalNoteShow)
		g.PUT("/criminal/:id", controllers.CriminalNoteEdit)
		g.DELETE("/criminal/:id", controllers.CriminalNoteDelete)

		// experience
		g.GET("/experiences", controllers.ExperienceIndex)
		g.POST("/experience", controllers.ExperienceCreate)
		g.GET("/experience/:id", controllers.ExperienceShow)
		g.PUT("/experience/:id", controllers.ExperienceEdit)
		g.DELETE("/experience/:id", controllers.ExperienceDelete)

		// reference
		g.GET("/references", controllers.ReferenceIndex)
		g.POST("/reference", controllers.ReferenceCreate)
		g.GET("/reference/:id", controllers.ReferenceShow)
		g.PUT("/reference/:id", controllers.ReferenceEdit)
		g.DELETE("/reference/:id", controllers.ReferenceDelete)
	}
}
