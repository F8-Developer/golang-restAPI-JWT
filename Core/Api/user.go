package Api

import (
	"golang.org/x/crypto/bcrypt"
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Models"
	"golang-restAPI-JWT/Auth"
)

// Register one new user in db
func RegisterUser(reg_req Structs.RegisterRequest) (reg_res Structs.RegisterResponse) {
	reg_res.ResponseCode = 200
	reg_res.ResponseMsg = "User successfully register"

	var user Models.User
	err := Models.FindUser(&user, reg_req.Email)
	if err == nil {
		reg_res.ResponseCode = 203
		reg_res.ResponseMsg = "User email already used"
		return
	}

	user.Name = reg_req.Name
	user.Email = reg_req.Email
	hashedPassword, err_pass := bcrypt.GenerateFromPassword([]byte(reg_req.Password), bcrypt.DefaultCost)
	if err_pass != nil {
		reg_res.ResponseCode = 204
		reg_res.ResponseMsg = err_pass.Error()
		return
	}
	user.Password = string(hashedPassword)

	err = Models.CreateUser(&user)
	if err != nil {
		reg_res.ResponseCode = 205
		reg_res.ResponseMsg = err.Error()
		return
	}
	return reg_res;
}

// Login user and generate jwt auth
func LoginUser(log_req Structs.LoginRequest) (log_res Structs.LoginResponse) {
	log_res.ResponseCode = 200
	log_res.ResponseMsg = "User successfully login"

	var user Models.User
	err := Models.FindUser(&user, log_req.Email)
	if err != nil {
		log_res.ResponseCode = 206
		log_res.ResponseMsg = "Failed to login, please check you email password"
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(log_req.Password))
	if err != nil {
		log_res.ResponseCode = 206
		log_res.ResponseMsg = "Failed to login, please check you email password"
		return
	}

	tokenString, err:= Auth.GenerateJWT(user.Name, user.Email)
	if err != nil {
		log_res.ResponseCode = 207
		log_res.ResponseMsg = "Invalid generate token"
		return
	}
	log_res.Token = tokenString
	return log_res;
}
