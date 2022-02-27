package VaOnline

import (
	"intrajasa-merchant-api-gateway/Core/Structs"
)

// Register register one new user in db, return a boolean value to make know success or not.
func GenerateVa(gv_req Structs.GenerateVaRequest) (gv_res Structs.GenerateVaResponse) {
	gv_res.MerchantId = gv_req.MerchantId
	gv_res.MerchantRefCode = gv_req.MerchantRefCode
	gv_res.TotalAmount = gv_req.TotalAmount
	gv_res.VaNumber = "8228006200100634"
	gv_res.ResponseCode = 200
	gv_res.ResponseMsg = "Success generate VA Number"

	return gv_res
}
