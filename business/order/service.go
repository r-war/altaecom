package order

import (
	"AltaEcom/business"
	"AltaEcom/util/validator"
	"time"
)

type OrderSpec struct {
	UserID int
	Status bool
}

type OrderItemSpec struct {
	ProductID int
	OrderID   int
	Price     int
	Qty       int
}

type service struct {
	repo          Repository
	repoOrderItem RepositoryOrderItem
}

func NewService(repo Repository, repoOrderItem RepositoryOrderItem) Service {
	return &service{repo, repoOrderItem}
}

func (s *service) GetOrderByUserID(userID int) (*Order, error) {
	return s.repo.GetOrderByUserID(userID)
}

func (s *service) NewOrderByUserID(userID int) (*Order, error) {
	return s.repo.NewOrderByUserID(userID, time.Now())
}

func (s *service) GetOrderItemByUserID(id int) (*OrderDetail, error) {
	cart, err := s.repo.GetOrderByUserID(id)

	if err != nil {
		return nil, err
	}

	allItems, err := s.repoOrderItem.GetOrderItemByOrderID(id)
	if err != nil {
		return nil, err
	}

	items := NewOrderDetail(cart, allItems)

	return items, nil
}

func (s *service) AddItemToOrder(orderID int, product *OrderItemSpec) error {
	if err := validator.GetValidator().Struct(product); err != nil {
		return business.ErrInvalidSpec
	}

	data := NewOrderItem(
		product.ProductID,
		product.OrderID,
		product.Price,
		product.Qty,
		time.Now(),
		time.Now())
	return s.repoOrderItem.AddItemToOrder(orderID, data)
}

func (s *service) UpdateItemInOrder(orderID int, product OrderItem) error {
	return s.repoOrderItem.UpdateItemInOrder(orderID, product)
}

func (s *service) RemoveItemInOrder(orderID int, productID int) error {
	return s.repoOrderItem.RemoveItemInOrder(orderID, productID)
}
