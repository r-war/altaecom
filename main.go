package main

import (
	"AltaEcom/api"
	productController "AltaEcom/api/product"
	productService "AltaEcom/business/product"
	"AltaEcom/config"
	"AltaEcom/modules/migration"
	productRepository "AltaEcom/modules/product"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection(conf *config.AppConfig) *gorm.DB {

	// config := map[string]string{
	// 	"DB_Username": "root",
	// 	"DB_Password": "",
	// 	"DB_Port":     "3306",
	// 	"DB_Host":     "127.0.0.1",
	// 	"DB_Name":     "altaecom",
	// }

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DbUsername,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName)

	db, e := gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db

}

func main() {
	cfg := config.GetConfig()

	dbConnect := DBConnection(cfg)

	productRepo := productRepository.NewRepository(dbConnect)

	productService := productService.NewService(productRepo)

	productController := productController.NewController(productService)

	e := echo.New()
	api.RegisterPath(
		e,
		productController,
	)

	func() {
		address := fmt.Sprintf(":%d", cfg.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("Shutdown Echo Service")
		}

	}()
}
