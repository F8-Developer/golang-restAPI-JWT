package Models

import (
	"golang-restAPI-JWT/Database"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	ID				uint		`gorm:"primaryKey"`
	Name			string		`gorm:"column:name"`
	Descriptions	string		`gorm:"column:descriptions"`
}

func (ctg *Category) TableName() string {
	return "categories"
}

// Get First Category by return error info.
// 	err := Models.FindCategory(&category, "category_id")
func FirstCategory(ctg *Category) error {
	err := Database.Mysql.First(ctg).Error
	return err
}

// Insert category which will be saved in database returning with error info
// 	if err := CreateCategory(&Category); err != nil { ... }
func CreateCategory(ctg *Category) error {
	err := Database.Mysql.Save(ctg).Error
	return err
}