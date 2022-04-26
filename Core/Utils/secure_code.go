package Utils

import (
	"fmt"
	"crypto/sha1"
	"crypto/sha256"
	"strconv"
	"golang-restAPI-JWT/Core/Models"
)

func SecureCodeCheck(secure_code string, merchant_refcode string, merchant_va Models.MerchantVa) bool {
	// secredword from merchant table, generate to sha1
	sha := sha1.New()
	sha.Write([]byte(merchant_va.SecretWord))
	encrypted := sha.Sum(nil)
    encryptedString := fmt.Sprintf("%x", encrypted)
	// =========
	string_merchant_id := strconv.FormatUint(uint64(merchant_va.ID), 10)
	sum := sha256.Sum256([]byte(string_merchant_id+merchant_refcode+encryptedString))
	sumString := fmt.Sprintf("%x", sum)

	if sumString != secure_code {
		return false
	}
	return true
}