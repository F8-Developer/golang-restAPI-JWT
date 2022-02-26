package structs

type GetTokenRequest struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	SecureCode string `json:"secureCode" validate:"required,alphanum,max=64"`
}

type GetTokenResponse struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	Token string `json:"token" validate:"required,alphanum,max=64"`
	ResponseCode int `json:"responseCode" validate:"required,numeric,max=999"`
	ResponseMsg string `json:"responseMsg" validate:"required,alphanum,max=255"`
}