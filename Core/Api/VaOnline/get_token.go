package VaOnline

import (
	// "fmt"
	"intrajasa-merchant-api-gateway/Core/Structs"
	"intrajasa-merchant-api-gateway/Core/Models"
)

// Register register one new user in db, return a boolean value to make know success or not.
func GenerateToken(gt_req Structs.GetTokenRequest) (gt_res Structs.GetTokenResponse) {
	gt_res.MerchantId = gt_req.MerchantId
	gt_res.MerchantRefCode = gt_req.MerchantRefCode
	gt_res.ResponseCode = 200
	gt_res.ResponseMsg = "Success generate token"
	
	var merchant_va Models.MerchantVa
	err := Models.FindMerchantVa(&merchant_va, gt_res.MerchantId)
	if err != nil {
		gt_res.ResponseCode = 400
		gt_res.ResponseMsg = "Merchant Not Found!"
		return gt_res
	}

	return gt_res
}
