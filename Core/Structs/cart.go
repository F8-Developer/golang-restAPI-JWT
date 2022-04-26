package Structs

type CartAddRequest struct {
	ProductID	uint	`json:"productID" validate:"required,uint-len-10"`
	Quantity	int	`json:"quantity" validate:"required,int-len-11"`
}

type CartAddResponse struct {
	ResponseCode	int				`json:"responseCode"`
	ResponseMsg		string			`json:"responseMsg"`
}