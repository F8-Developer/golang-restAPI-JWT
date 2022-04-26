package Models

import (
	"golang-restAPI-JWT/Database"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID				uint	`json:"-";gorm:"primaryKey"`
	Name			string	`gorm:"column:name"`
	Email			string	`gorm:"column:email;unique_index"`
	Password		string	`json:"-";gorm:"column:password"`
	RememberToken	string	`json:"-";gorm:"column:remember_token"`
}

func (usr *User) TableName() string {
	return "users"
}

// Find user by email address and return error info.
// 	err := Models.FindUser(&user, "user_email_address")
func FindUser(usr *User, email string) error {
	err := Database.Mysql.Where("email = ?", email).First(usr).Error
	return err
}

// Insert user which will be saved in database returning with error info
// 	if err := CreateUser(&user); err != nil { ... }
func CreateUser(usr *User) error {
	err := Database.Mysql.Save(usr).Error
	return err
}