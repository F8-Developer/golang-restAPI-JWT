package vaonline

import (
	"intrajasa-merchant-api-gateway/Core/Structs"
)

// Register register one new user in db, return a boolean value to make know success or not.
func UpdateVa(uv_req structs.UpdateVaRequest) (uv_res structs.UpdateVaResponse) {
	uv_res.ResponseCode = 200
	uv_res.ResponseMsg = "Success"

	return uv_res
}
