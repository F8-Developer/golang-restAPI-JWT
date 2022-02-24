package vaonline

import (
)

// Register register one new user in db, return a boolean value to make know success or not.
func (gts GetTokenService) GenerateToken() (getTokenInfo GetTokenService) {
	gts.merchantId = "001"
	gts.merchantRefCode = "JS008sKs"
	gts.token = "RTkwQjk2QzVGQUM4NDIwQzYxMDVCNDI4QUFCNTNGRkEwRkJCNDBEODA4NEIxOUQ1MTc1NjcyMTFGNDBCNUVBOQ=="
	gts.responseCode = "200"
	gts.responseMsg = "Success generate token"

	return getTokenInfo
}
