package routes

import (
	"crud-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/admin-api")
	{
		grpl.GET("admin", controllers.GetAdmins)
		grpl.POST("admin", controllers.CreateAdmin)
		grpl.GET("admin/:username", controllers.GetAdminByUsername)
		grpl.PUT("admin/:username", controllers.UpdateAdmin)
		grpl.DELETE("admin/:username", controllers.DeleteAdmin)
	}
	return r
}
