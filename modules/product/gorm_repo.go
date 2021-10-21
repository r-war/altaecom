package product

import (
	"AltaEcom/business/product"
	"AltaEcom/modules/category"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type ProductTable struct {
	ID          int                    `gorm:"id"`
	CategoryId  int                    `gorm:"category_id"`
	Name        string                 `gorm:"name"`
	Price       int                    `gorm:"price"`
	Qty         int                    `gorm:"qty"`
	Description string                 `gorm:"description"`
	Image       string                 `gorm:"image"`
	CreatedAt   time.Time              `gorm:"created_at"`
	UpdatedAt   time.Time              `gorm:"updated_at"`
	DeletedAt   time.Time              `gorm:"DeletedAt"`
	Category    category.CategoryTable `gorm:"foreignkey:CategoryId"`
}

func newProductTable(product product.Product) *ProductTable {
	return &ProductTable{
		product.ID,
		product.CategoryId,
		product.Name,
		product.Price,
		product.Qty,
		product.Description,
		product.Image,
		product.CreatedAt,
		product.UpdatedAt,
		product.DeletedAt,
		category.CategoryTable{},
	}
}

func (col *ProductTable) ToProduct() product.Product {
	var product product.Product

	product.ID = col.ID
	product.CategoryId = col.CategoryId
	product.Name = col.Name
	product.Price = col.Price
	product.Qty = col.Qty
	product.Description = col.Description
	product.Image = col.Image
	product.CreatedAt = col.CreatedAt
	product.UpdatedAt = col.UpdatedAt
	product.DeletedAt = col.DeletedAt

	return product
}

func newGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (repo *GormRepository) findProductByID(id int) (*product.Product, error) {
	var data ProductTable

	err := repo.DB.Where("id = ?", id).First(&data).Error

	if err != nil {
		return nil, err
	}

	product := data.ToProduct()

	return &product, nil
}
