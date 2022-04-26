package Api

import (
	"golang.org/x/crypto/bcrypt"
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Models"
)

// Register one new user in db, return a boolean value to make know success or not.
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
		reg_res.ResponseCode = 204
		reg_res.ResponseMsg = err.Error()
		return
	}
	return reg_res;
}
