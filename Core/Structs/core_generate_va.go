package Structs

type CoreGenerateVaResponse struct {
	Data CoreDataGVResponse `json:"data"`
	Message string `json:"message"`
	Status int `json:"status"`
}

type CoreDataGVResponse struct {
	VaNo string `json:"va_no"`
    Message string `json:"message"`
    StatusCode string `json:"status_code"`
}