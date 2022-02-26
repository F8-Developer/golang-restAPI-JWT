package structs

type GetVaPaymentStatusRequest struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	VaNumber string `json:"vaNumber" validate:"required,numeric,max=200"`
	TotalAmount float64 `json:"totalAmount" validate:"required,float-decimal-val"`
	SecureCode string `json:"secureCode" validate:"required,alphanum,max=64"`
}

type GetVaPaymentStatusResponse struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	VaNumber string `json:"vaNumber" validate:"required,numeric,max=200"`
	TotalAmount float64 `json:"totalAmount" validate:"required,float-decimal-val"`
	PaymentStatus int `json:"paymentStatus" validate:"required,int-length-1"`
	ResponseCode int `json:"responseCode" validate:"required,numeric,max=999"`
	ResponseMsg string `json:"responseMsg" validate:"required,alphanum,max=255"`
}