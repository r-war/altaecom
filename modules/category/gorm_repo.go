package category

import (
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type CategoryTable struct {
	gorm.Model
	Name string `gorm:"name"`
}
