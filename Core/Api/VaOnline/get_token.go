package vaonline

import (
	"fmt"
	"intrajasa-merchant-api-gateway/Core/Structs"
	"intrajasa-merchant-api-gateway/Core/Models"
)

// Register register one new user in db, return a boolean value to make know success or not.
func GenerateToken(gt_req structs.GetTokenRequest) (gt_res structs.GetTokenResponse) {
	gt_res.MerchantId = gt_req.MerchantId
	gt_res.MerchantRefCode = gt_req.MerchantRefCode
	gt_res.Token = "RTkwQjk2QzVGQUM4NDIwQzYxMDVCNDI4QUFCNTNGRkEwRkJCNDBEODA4NEIxOUQ1MTc1NjcyMTFGNDBCNUVBOQ=="
	gt_res.ResponseCode = 200
	gt_res.ResponseMsg = "Success generate token"
	
	var merchant_va Models.MerchantVa
	err := Models.FindMerchantVa(&merchant_va, gt_res.MerchantId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(merchant_va.SecretWord)

	return gt_res
}
