//Models/Model.go

package Models

import (
	"crud-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllAdmins Fetch all admin data
func GetAllAdmins(admin *[]Admin) (err error) {
	if err = Config.DB.Find(admin).Error; err != nil {
		return err
	}
	return nil
}

//CreateAdmin ... Insert New data
func CreateAdmin(admin *Admin) (err error) {
	if err = Config.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

//GetAdminByUsername ... Fetch only one admin by Username
func GetAdminByUsername(admin *Admin, username string) (err error) {
	if err = Config.DB.Where("username = ?", username).First(admin).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAdmin ... Update admin
func UpdateAdmin(admin *Admin, id string) (err error) {
	fmt.Println(admin)
	Config.DB.Save(admin)
	return nil
}

//DeleteAdmin ... Delete admin
func DeleteAdmin(admin *Admin, username string) (err error) {
	Config.DB.Where("username = ?", username).Delete(admin)
	return nil
}
