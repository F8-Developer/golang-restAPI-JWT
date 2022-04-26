package Api

import (
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Models"
)

func GetAllCategories() (cas_res Structs.CategoriesResponse) {
	cas_res.ResponseCode = 200
	cas_res.ResponseMsg = "List Categories successful"
	var categories []Models.Category
	err := Models.FindAllCategories(&categories)
	if err != nil {
		cas_res.ResponseCode = 205
		cas_res.ResponseMsg = err.Error()
	}

	cas_res.Categories = categories
	return cas_res;
}