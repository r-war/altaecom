package migration

import (
	product "AltaEcom/modules/product"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&product.ProductTable{})
}
