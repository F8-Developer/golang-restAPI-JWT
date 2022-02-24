package structs

type GenerateVaRequest struct {
	MerchantId string `json:"merchantId" validate:"required,alphanum,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,alphanum,max=20"`
	CustomerData GenerateVaCustomerData `json:"customerData" validate:"required"`
	ProductData []GenerateVaProductData `json:"productData"`
    TotalAmount float64 `json:"totalAmount" validate:"required"`
    VaType int `json:"vaType" validate:"required,len=1"`
    ExpiryPeriod int `json:"expiryPeriod" validate:"numeric"`
	SecureCode string `json:"secureCode" validate:"required,is-awesome,alphanum,max=64"`
}

type GenerateVaResponse struct {
	MerchantId string `json:"merchantId" validate:"required,max=12"`
    MerchantRefCode string `json:"merchantRefCode" validate:"required,max=20"`
	CustomerData GenerateVaCustomerData `json:"customerData"`
	ProductData []GenerateVaProductData `json:"productData"`
    TotalAmount string `json:"totalAmount"`
    VaNumber string `json:"vaNumber"`
    VaExpiryDate string `json:"vaExpiryDate"`
	ResponseCode int `json:"responseCode"`
	ResponseMsg string `json:"responseMsg"`
}

type GenerateVaCustomerData struct {
    CustName string `json:"custName" validate:"required,max=100"`
    CustAddress1 string `json:"custAddress1"`
    CustAddress2 string `json:"custAddress2"`
    CustAddress3 string `json:"custAddress3"`
    CustPhoneNumber int `json:"custPhoneNumber"`
    CustEmail string `json:"custEmail"`
    CustRegisteredDate string `json:"custRegisteredDate"`
    CustCountryCode string `json:"custCountryCode"`
}

type GenerateVaProductData struct {
    ProductCode string `json:"productCode" validate:"max=100"`
    ProductName string `json:"productName" validate:"max=100"`
    ProductQuantity int `json:"productQuantity" validate:"max=5"`
    ProductAmount string `json:"productAmount" validate:"max=100"`
}