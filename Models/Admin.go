//Models/Model.go

package models

import (
	"crud-api/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllAdmins Fetch all admin data
func GetAllAdmins(admin *[]Admin) (err error) {
	if err = config.DB.Find(admin).Error; err != nil {
		return err
	}
	return nil
}

//CreateAdmin ... Insert New data
func CreateAdmin(admin *Admin) (err error) {
	if err = config.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

//GetAdminByUsername ... Fetch only one admin by Username
func GetAdminByUsername(admin *Admin, username string) (err error) {
	if err = config.DB.Where("username = ?", username).First(admin).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAdmin ... Update admin
func UpdateAdmin(admin *Admin, username string) (err error) {
	fmt.Println(admin)
	config.DB.Save(admin)
	return nil
}

//DeleteAdmin ... Delete admin
func DeleteAdmin(admin *Admin, username string) (err error) {
	config.DB.Where("username = ?", username).Delete(admin)
	return nil
}
