package Utils

import (
	"fmt"
	"time"
	"strconv"
	"encoding/base64"
	"crypto/sha256"
	"intrajasa-merchant-api-gateway/Core/Models"
	"intrajasa-merchant-api-gateway/Core/Utils/Redis"
)

func GenerateToken(merchant_refcode string, merchant_va Models.MerchantVa) (string, string) {
	t := time.Now()
	tUnix := t.Unix()
	stringTimestamp := strconv.FormatInt(tUnix, 10)
	string_merchant_id := strconv.FormatUint(uint64(merchant_va.ID), 10)

	string_token_base64 := base64.StdEncoding.EncodeToString([]byte(string_merchant_id+merchant_va.SecretWord+stringTimestamp))
	sum := sha256.Sum256([]byte(string_token_base64))
	string_token_sha256 := fmt.Sprintf("%x", sum)
	return string_token_base64, string_token_sha256
}

func ValidateToken(t string) bool {
	val, err := Redis.Client.Get("t"+t).Result()
	if err != nil {
		return false
    }

	if val == t {
		err = Redis.Client.Del("t"+t).Err()
		if err != nil {
			return false
		} else {
			return true
		}
	}
	return false
}