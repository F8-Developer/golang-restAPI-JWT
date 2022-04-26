package Api

import (
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Models"
)

func GetCart(email string) (cl_res Structs.CartListResponse) {
	cl_res.ResponseCode = 200
	cl_res.ResponseMsg = "List Cart successful"
	var user Models.User
	var carts []Models.Cart
	err := Models.FindUser(&user, email)
	if err != nil {
		cl_res.ResponseCode = 206
		cl_res.ResponseMsg = "Failed to find user"
		return
	}
	err = Models.FindCartByUserID(user.ID, &carts)
	if err != nil {
		cl_res.ResponseCode = 205
		cl_res.ResponseMsg = err.Error()
	}

	cl_res.Cart = carts
	return cl_res;
}

func AddToCart(email string, ca_req Structs.CartAddRequest) (ca_res Structs.CartAddResponse) {
	ca_res.ResponseCode = 200
	ca_res.ResponseMsg = "Product successfully add to cart"
	var user Models.User
	var product Models.Product
	var cart Models.Cart

	err := Models.FindUser(&user, email)
	if err != nil {
		ca_res.ResponseCode = 206
		ca_res.ResponseMsg = "Failed to find user"
		return
	}
	err = Models.FindProduct(&product, ca_req.ProductID)
	if err != nil {
		ca_res.ResponseCode = 206
		ca_res.ResponseMsg = "Failed to find product"
		return
	}
	if ca_req.Quantity > product.Quantity {
		ca_res.ResponseCode = 207
		ca_res.ResponseMsg = "Quantity more than product stock"
		return
	}

	// add to cart
	cart.UserID = user.ID
	cart.ProductID = product.ID
	cart.Quantity = ca_req.Quantity
	cart.Price = product.Price
	cart.Total = float64(ca_req.Quantity) * product.Price
	err = Models.CartAdd(&cart)
	if err != nil {
		ca_res.ResponseCode = 206
		ca_res.ResponseMsg = "Failed to add cart"
		return
	}
	// decrease product quantity
	err = Models.UpdateProductQuantity(&product, ca_req.Quantity * -1)
	return ca_res;
}