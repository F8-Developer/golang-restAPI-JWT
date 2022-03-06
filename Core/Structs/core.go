package Structs

type CoreResponse struct {
	Data CoreDataResponse `json:"data"`
	Message string `json:"message"`
	Status int `json:"status"`
}

type CoreDataResponse struct {
	VaNo string `json:"va_no"`
    Message string `json:"message"`
    StatusCode string `json:"status_code"`
}