package Models

import (
	"golang-restAPI-JWT/Database"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID				uint		`gorm:"primaryKey"`
	Categories		[]Category	`gorm:"many2many:products_categories;"`
	Name			string		`gorm:"column:name"`
	Descriptions	string		`gorm:"column:descriptions"`
	Quantity		int			`gorm:"column:quantity"`
	Price			float64		`gorm:"column:price;type:decimal(18,2);"`
}

func (prd *Product) TableName() string {
	return "products"
}

func FindAllProducts(prd *[]Product) error {
	err := Database.Mysql.Preload("Categories").Find(&prd).Error
	return err
}

func FindProduct(prd *Product, id uint) error {
	err := Database.Mysql.Preload("Categories").First(&prd, id).Error
	return err
}

// Get First Product by return error info.
// 	err := Models.FindProduct(&product, "product_id")
func FirstProduct(prd *Product) error {
	err := Database.Mysql.First(prd).Error
	return err
}

// Insert product which will be saved in database returning with error info
// 	if err := CreateProduct(&Product); err != nil { ... }
func CreateProduct(prd *Product, ctg *Category) error {
	prd.Categories = append(prd.Categories, Category{ID: ctg.ID, Name: ctg.Name, Descriptions: ctg.Name})
	err := Database.Mysql.Create(&prd).Error
	if err != nil {
		return err
	}
	return err
}