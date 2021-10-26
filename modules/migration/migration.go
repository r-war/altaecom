package migration

import (
	admin "AltaEcom/modules/admin"
	category "AltaEcom/modules/category"
	product "AltaEcom/modules/product"
	user "AltaEcom/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&product.Product{}, 
		&category.Category{},
		&admin.Admin{},
		&user.User{})
}
