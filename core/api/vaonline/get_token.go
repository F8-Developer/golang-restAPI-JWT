package vaonline

import (
	"intrajasa-merchant-api-gateway/core/structs"
)

// Register register one new user in db, return a boolean value to make know success or not.
func GenerateToken(gt_req structs.GetTokenRequest) (gt_res structs.GetTokenResponse) {
	gt_res.MerchantId = gt_req.MerchantId
	gt_res.MerchantRefCode = gt_req.MerchantRefCode
	gt_res.Token = "RTkwQjk2QzVGQUM4NDIwQzYxMDVCNDI4QUFCNTNGRkEwRkJCNDBEODA4NEIxOUQ1MTc1NjcyMTFGNDBCNUVBOQ=="
	gt_res.ResponseCode = 200
	gt_res.ResponseMsg = "Success generate token"

	return gt_res
}
