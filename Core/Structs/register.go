package Structs

type RegisterRequest struct {
	Name		string `json:"name" validate:"required,max=255"`
    Email		string `json:"email" validate:"required,email,max=255"`
	Password	string `json:"password" validate:"required,max=100"`
}

type RegisterResponse struct {
	ResponseCode	int		`json:"responseCode"`
	ResponseMsg		string	`json:"responseMsg"`
}