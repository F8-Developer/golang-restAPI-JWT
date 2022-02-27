package Structs

type UpdateVaRequest struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	VaNumber string `json:"vaNumber" validate:"required,numeric,max=200"`
	Amount float64 `json:"amount" validate:"required,float-decimal-val"`
	SecureCode string `json:"secureCode" validate:"required,alphanum,max=64"`
}

type UpdateVaResponse struct {
	ResponseCode int `json:"responseCode" validate:"required,numeric,max=999"`
	ResponseMsg string `json:"responseMsg" validate:"required,alphanum,max=255"`
}