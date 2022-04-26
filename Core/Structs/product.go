package Structs

import (
	"golang-restAPI-JWT/Core/Models"
)

type ProductsResponse struct {
	Products		[]Models.Product	`json:"products"`
	ResponseCode	int					`json:"responseCode"`
	ResponseMsg		string				`json:"responseMsg"`
}