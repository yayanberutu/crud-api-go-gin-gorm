package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//DBConfig Untuk merepresentasikan database
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		"localhost",
		3306,
		"root",
		"",
		"del",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
