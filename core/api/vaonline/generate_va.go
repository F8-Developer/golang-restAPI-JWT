package vaonline

import (
	"intrajasa-merchant-api-gateway/core/structs"
)

// Register register one new user in db, return a boolean value to make know success or not.
func GenerateVa(gv_req structs.GenerateVaRequest) (gv_res structs.GenerateVaResponse) {
	gv_res.MerchantId = gv_req.MerchantId
	gv_res.MerchantRefCode = gv_req.MerchantRefCode
	gv_res.TotalAmount = gv_req.TotalAmount
	gv_res.VaNumber = "8228006200100634"
	gv_res.ResponseCode = 200
	gv_res.ResponseMsg = "Success generate VA Number"

	return gv_res
}
