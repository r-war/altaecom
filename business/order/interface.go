package order

import "time"

type Service interface {
	GetOrderByUserID(userID int) (*Order, error)

	GetOrderItemByUserID(id int) (*OrderDetail, error)

	NewOrderByUserID(userID int) (*Order, error)

	AddItemToOrder(orderID int, product *OrderItemSpec) error

	UpdateItemInOrder(orderID int, product OrderItem) error

	RemoveItemInOrder(OrderID int, productID int) error
}

type Repository interface {
	GetOrderByUserID(userID int) (*Order, error)

	NewOrderByUserID(userID int, createdAt time.Time) (*Order, error)
}

type RepositoryOrderItem interface {
	GetOrderItemByOrderID(orderID int) (*[]OrderItem, error)

	AddItemToOrder(orderID int, product OrderItem) error

	UpdateItemInOrder(orderID int, product OrderItem) error

	RemoveItemInOrder(orderID int, productID int) error
}
