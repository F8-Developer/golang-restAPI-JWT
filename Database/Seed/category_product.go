package Seed

import (
	"golang-restAPI-JWT/Core/Models"
)

func CategoryProductSeed() error {
	var category Models.Category
	var product Models.Product

	// =======
	category.Name = "buah"
	category.Descriptions = "jenis-jenis buahan"
	err := Models.CreateCategory(&category)
	// ==
	product.Name = "Anggur Merah Globe"
	product.Descriptions = "Anggur Merah Globe (~0.65 kg)"
	product.Quantity = 100
	product.Price = 56200.00
	err = Models.CreateProduct(&product,&category)
	product = Models.Product{}
	// ==
	product.Name = "Apel Fuji Jingle"
	product.Descriptions = "Apel Fuji Jingle (~0.3 kg)"
	product.Quantity = 100
	product.Price = 15000.00
	err = Models.CreateProduct(&product,&category)
	product = Models.Product{}
	// ==
	category = Models.Category{}

	// =======
	category.Name = "sayur"
	category.Descriptions = "jenis-jenis sayuran"
	err = Models.CreateCategory(&category)
	// ==
	product.Name = "Wortel"
	product.Descriptions = "Wortel (~0.1 kg)"
	product.Quantity = 100
	product.Price = 1700.00
	err = Models.CreateProduct(&product,&category)
	product = Models.Product{}
	// ==
	product.Name = "Kol Putih"
	product.Descriptions = "Kol Putih (~1 kg)"
	product.Quantity = 100
	product.Price = 18000.00
	err = Models.CreateProduct(&product,&category)
	product = Models.Product{}
	// ==
	category = Models.Category{}

	// =======
	category.Name = "rempah-rempah"
	category.Descriptions = "jenis-jenis rempah-rempah"
	err = Models.CreateCategory(&category)
	// ==
	product.Name = "Lengkuas"
	product.Descriptions = "Lengkuas (~0.2 kg)"
	product.Quantity = 100
	product.Price = 8000.00
	err = Models.CreateProduct(&product,&category)
	product = Models.Product{}
	// ==
	product.Name = "Jahe"
	product.Descriptions = "Jahe (~0.2 kg)"
	product.Quantity = 100
	product.Price = 10800.00
	err = Models.CreateProduct(&product,&category)
	product = Models.Product{}
	// ==
	category = Models.Category{}

	return err
}
