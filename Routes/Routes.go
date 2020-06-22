package Routes

import (
	"crud-api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grpl := r.Group("/admin-api")
	{
		grpl.GET("admin", Controllers.GetAdmins)
		grpl.POST("admin", Controllers.CreateAdmin)
		grpl.GET("admin/:username", Controllers.GetAdminByUsername)
		grpl.PUT("admin/:username", Controllers.UpdateAdmin)
		grpl.DELETE("admin/:username", Controllers.DeleteAdmin)
	}
	return r
}
