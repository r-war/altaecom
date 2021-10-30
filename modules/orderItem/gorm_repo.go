package orderitem

import (
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type OrderItem struct {
	ID        int `gorm:"id;primaryKey;autoIncrement"`
	OrderID   int `gorm:"order_id;"`
	ProductID int `gorm:"product_id"`
	Qty       int `gorm:"qty"`
	Price     int `gorm:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItemWithProductName struct {
	OrderItem
	ProductName string `gorm:"name"`
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}
func NewOrderItemTable(
	cartID int,
	item OrderItem,
) OrderItem {
	return OrderItem{
		OrderID:   cartID,
		ProductID: item.ProductID,
		Qty:       item.Qty,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func (r *Repository) AddItemToOrder(orderID int, item OrderItem) error {
	err := r.DB.Create(NewOrderItemTable(orderID, item))

	if err != nil {
		return gorm.ErrDryRunModeUnsupported
	}

	return nil
}

func (r *Repository) GetOrderItemByUserID(OrderID int) (*[]OrderItemWithProductName, error) {
	var itemsDetail *[]OrderItemWithProductName

	err := r.DB.Raw("select t1.*, t2.name product_name from order_item t1 inner join products t2 on t2.id = t1.product_id where t1.order_id = ?", OrderID).Scan(itemsDetail).Error

	if err != nil {
		return nil, err
	}
	return itemsDetail, nil
}

func (r *Repository) UpdateItemInOrder(orderID int, item OrderItem) error {
	var orderItem OrderItem

	err := r.DB.Where("order_id = ? and id = ?", orderID, item.ID).Find(orderItem).Error
	if err != nil {
		return err
	}

	r.DB.Model(orderItem).Updates(OrderItem{
		Price:     item.Price,
		Qty:       item.Qty,
		UpdatedAt: time.Now(),
	})

	return nil
}

func (r *Repository) RemoveItemInOrder(orderID int, productID int) error {
	var orderItem OrderItem

	err := r.DB.Where("order_id = ? and product_id = ? ", orderID, productID).Find(orderItem).Error

	if err != nil {
		return err
	}

	r.DB.Delete(orderItem)
	return nil
}
