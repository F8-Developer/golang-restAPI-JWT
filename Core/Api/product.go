package Api

import (
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Models"
)

func GetAllProducts() (prs_res Structs.ProductsResponse) {
	prs_res.ResponseCode = 200
	prs_res.ResponseMsg = "List Products successful"
	var products []Models.Product
	err := Models.FindAllProducts(&products)
	if err != nil {
		prs_res.ResponseCode = 205
		prs_res.ResponseMsg = err.Error()
	}

	prs_res.Products = products
	return prs_res;
}

func GetProductDetail(prd_req Structs.ProductDetailRequest) (prd_res Structs.ProductDetailResponse) {
	prd_res.ResponseCode = 200
	prd_res.ResponseMsg = "Product detail successful"
	var product Models.Product
	err := Models.FindProduct(&product, prd_req.ProductID)
	if err != nil {
		prd_res.ResponseCode = 205
		prd_res.ResponseMsg = err.Error()
	}

	prd_res.Product = product
	return prd_res;
}