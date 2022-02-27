package Structs

type GenerateVaRequest struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	CustomerData GenerateVaCustomerData `json:"customerData" validate:"required"`
	ProductData []GenerateVaProductData `json:"productData" validate:"dive"`
    TotalAmount float64 `json:"totalAmount" validate:"required,float-decimal-val"`
    VaType *int `json:"vaType" validate:"required,int-length-1"`
    ExpiryPeriod int `json:"expiryPeriod" validate:"numeric,int-length-10"`
	SecureCode string `json:"secureCode" validate:"required,alphanum,max=64"`
}

type GenerateVaResponse struct {
	MerchantId string `json:"merchantId" validate:"required,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,max=20"`
    TotalAmount float64 `json:"totalAmount" validate:"required,float-decimal-val"`
    VaNumber string `json:"vaNumber" validate:"required,numeric"`
	ResponseCode int `json:"responseCode" validate:"required,max=999"`
	ResponseMsg string `json:"responseMsg" validate:"required,max=255"`
}

type GenerateVaCustomerData struct {
    CustName string `json:"custName" validate:"required,max=100"`
    CustAddress1 string `json:"custAddress1" validate:"omitempty,max=100"`
    CustAddress2 string `json:"custAddress2" validate:"omitempty,max=100"`
    CustAddress3 string `json:"custAddress3" validate:"omitempty,max=100"`
    CustPhoneNumber *uint64 `json:"custPhoneNumber" validate:"numeric,uint64-length-18"`
    CustEmail string `json:"custEmail" validate:"omitempty,email"`
    CustRegisteredDate string `json:"custRegisteredDate" validate:"omitempty,date-format"`
    CustCountryCode string `json:"custCountryCode" validate:"omitempty,alphanum,max=3"`
}

type GenerateVaProductData struct {
    ProductCode string `json:"productCode" validate:"omitempty,max=100"`
    ProductName string `json:"productName" validate:"omitempty,max=100"`
    ProductQuantity int `json:"productQuantity" validate:"min=1,int-length-5"`
    ProductAmount float64 `json:"productAmount" validate:"float-decimal-val"`
}