package VaOnline

import (
	"fmt"
	"time"
	"intrajasa-merchant-api-gateway/Core/Structs"
	"intrajasa-merchant-api-gateway/Core/Models"
	"intrajasa-merchant-api-gateway/Core/Utils"
	"intrajasa-merchant-api-gateway/Core/Utils/Redis"
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
		gt_res.ResponseCode = 211
		gt_res.ResponseMsg = "Invalid Merchant Id"
		return gt_res
	}

	if !Utils.SecureCodeCheck(gt_req.SecureCode, gt_req.MerchantRefCode, merchant_va) {
		gt_res.ResponseCode = 213
		gt_res.ResponseMsg = "Invalid Secure Code"
		return
	}

	string_token_base64, string_token_sha256 := Utils.GenerateToken(gt_req.MerchantRefCode, merchant_va)
	gt_res.Token = string_token_base64

	// store to redis db
	err = Redis.Client.Set("t"+string_token_sha256, string_token_sha256, 10 * time.Minute).Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string_token_sha256)
	return gt_res
}
