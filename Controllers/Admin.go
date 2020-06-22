package Controllers

import (
	"crud-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAdmins ... Get all admins
func GetAdmins(c *gin.Context) {
	var admin []Models.Admin
	err := Models.GetAllAdmins(&admin)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

//CreateAdmin ... Create Admin
func CreateAdmin(c *gin.Context) {
	var admin Models.Admin
	c.BindJSON(&admin)
	err := Models.CreateAdmin(&admin)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}

}

//GetAdminByUsername ... Get the admin by username
func GetAdminByUsername(c *gin.Context) {
	username := c.Params.ByName("username")
	var admin Models.Admin
	err := Models.GetAdminByUsername(&admin, username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

//UpdateAdmin ... Update the admin information
func UpdateAdmin(c *gin.Context) {
	var admin Models.Admin
	username := c.Params.ByName("username")
	err := Models.GetAdminByUsername(&admin, username)
	if err != nil {
		c.JSON(http.StatusNotFound, admin)
	}
	c.BindJSON(&admin)
	err = Models.UpdateAdmin(&admin, username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, admin)
	}
}

//DeleteAdmin ... Delete the admin
func DeleteAdmin(c *gin.Context) {
	var admin Models.Admin
	username := c.Params.ByName("username")
	err := Models.DeleteAdmin(&admin, username)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"username" + username: "is deleted"})
	}
}
