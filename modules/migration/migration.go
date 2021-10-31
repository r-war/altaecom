package migration

import (
	admin "AltaEcom/modules/admin"
	category "AltaEcom/modules/category"
	order "AltaEcom/modules/order"
	orderitem "AltaEcom/modules/orderitem"
	product "AltaEcom/modules/product"
	user "AltaEcom/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&product.Product{},
		&category.Category{},
		&admin.Admin{},
		&user.User{},
		&order.Order{},
		&orderitem.OrderItem{},
	)
}
