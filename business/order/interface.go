package order

type Service interface {
	GetOrderByUserID(id int) (Order, error)

	GetOrderItemByUserID(id int) (OrderDetail, error)

	NewOrder(userID int) (Order, error)

	AddItemToOrder(orderID int, product Item) error

	UpdateItemInOrder(orderID int, product Item) error

	RemoveItemInOrder(OrderID int, productID int) error
}

type Repository interface {
	GetOrderByUserID(id int) (Order, error)

	NewOrder(userID int) (Order, error)
}

type RepositoryOrderItem interface {
	GetOrderItemByUserID(id int) (OrderDetail, error)

	AddItemToOrder(orderID int, product Item) error

	UpdateItemInOrder(orderID int, product Item) error

	RemoveItemInOrder(orderID int, productID int) error
}