package Models

import (
	"intrajasa-merchant-api-gateway/Database"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type MerchantVa struct {
	gorm.Model
	Id int64 `json:"id"`
	IdMerchant int64 `json:"id_merchant"`
	SecretWord string `json:"secret_word"`
}

func (mv *MerchantVa) TableName() string {
	return "merchant_va"
}

func GetAllMerchantVa(mv *[]MerchantVa) (err error) {
	if err = database.Mysql.Find(mv).Error; err != nil {
		return err
	}
	return nil
}

func FindMerchantVa(mv *MerchantVa, id string) (err error) {
	if err := database.Mysql.Where("id = ?", id).First(mv).Error; err != nil {
		return err
	}
	return nil
}