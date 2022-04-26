package Structs

import (
	"golang-restAPI-JWT/Core/Models"
)

type CartAddRequest struct {
	ProductID	uint	`json:"productID" validate:"required,uint-len-10"`
	Quantity	int		`json:"quantity" validate:"required,int-len-11"`
}

type CartAddResponse struct {
	ResponseCode	int		`json:"responseCode"`
	ResponseMsg		string	`json:"responseMsg"`
}

type CartListResponse struct {
	Cart			[]Models.Cart	`json:"cart"`
	ResponseCode	int				`json:"responseCode"`
	ResponseMsg		string			`json:"responseMsg"`
}