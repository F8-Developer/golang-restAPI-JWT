package Structs

type LoginRequest struct {
    Email		string `json:"email" validate:"required,email,max=255"`
	Password	string `json:"password" validate:"required,max=100"`
}

type LoginResponse struct {
	Token			string	`json:"token"`
	ResponseCode	int		`json:"responseCode"`
	ResponseMsg		string	`json:"responseMsg"`
}