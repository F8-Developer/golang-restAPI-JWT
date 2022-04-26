package Models

import (
	"golang-restAPI-JWT/Database"
	_ "github.com/go-sql-driver/mysql"
)

type Cart struct {
	ID			uint	`gorm:"primaryKey"`
	UserID		uint	`gorm:"column:user_id"`
	ProductID	uint	`gorm:"column:product_id"`
	User		Product	`gorm:"foreignKey:UserID"`
	Product		Product	`gorm:"foreignKey:ProductID"`
	Quantity	int		`gorm:"column:quantity"`
	Price		float64	`gorm:"column:price;type:decimal(18,2);"`
	Total		float64	`gorm:"column:total;type:decimal(18,2);"`
}

func (crt *Cart) TableName() string {
	return "carts"
}

// Insert user and product to cart which will be saved in database returning with error info
// 	if err := CreateCart(&Cart); err != nil { ... }
func CartAdd(crt *Cart) error {
	err := Database.Mysql.Save(crt).Error
	return err
}