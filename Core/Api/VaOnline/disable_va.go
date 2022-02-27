package VaOnline

import (
	"intrajasa-merchant-api-gateway/Core/Structs"
)

// Register register one new user in db, return a boolean value to make know success or not.
func DisableVa(dv_req Structs.DisableVaRequest) (dv_res Structs.DisableVaResponse) {
	dv_res.ResponseCode = 200
	dv_res.ResponseMsg = "Success"

	return dv_res
}
