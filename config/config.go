package config

import (
	"AltaEcom/modules/migration"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "Altaecom",
	}

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_name"])

	db, e := gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db

}
