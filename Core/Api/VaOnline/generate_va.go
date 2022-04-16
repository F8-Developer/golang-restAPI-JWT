package VaOnline

import (
	"merchant-api-gateway/Core/Structs"
	"merchant-api-gateway/Config"
	"io/ioutil"
	"net/http"
	"encoding/json"
    "time"
	"strconv"
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

var core_res Structs.CoreResponse

// Register register one new user in db, return a boolean value to make know success or not.
func GenerateVa(gv_req Structs.GenerateVaRequest) (gv_res Structs.GenerateVaResponse) {
	gv_res.MerchantId = gv_req.MerchantId
	gv_res.MerchantRefCode = gv_req.MerchantRefCode
	gv_res.TotalAmount = gv_req.TotalAmount

	core_secret_key := Config.GoDotEnvVariable("CORE_SECRET_KEY")
	url := Config.GoDotEnvVariable("CORE_URL_GENERATE_VA")

	// send request to core
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	va_type := strconv.Itoa(gv_req.VaType)
	amount := fmt.Sprintf("%f", gv_req.TotalAmount)
	expired_period := strconv.Itoa(gv_req.ExpiryPeriod)
	customer_phone := strconv.FormatUint(gv_req.CustomerData.CustPhoneNumber, 10)
	body_req, _ := json.Marshal(map[string]string{
        "merchant_id" : gv_req.MerchantId,
		"merchant_ref_code": gv_req.MerchantRefCode,
		"va_type" : va_type,
		"amount" : amount,
		"customer_name": gv_req.CustomerData.CustName,
		"customer_address": gv_req.CustomerData.CustAddress1,
		"customer_phone": customer_phone,
		"customer_email":gv_req.CustomerData.CustEmail,
		"expired_period": expired_period,
    })
	req, err := http.NewRequest("POST", url, bytes.NewReader(body_req))
	if err != nil {
        log.Fatal(err)
    }

    req.Header.Add("Core-Secret-Key", core_secret_key)
    resp, err := c.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

	json.Unmarshal([]byte(body), &core_res)
	gv_res.ResponseMsg = core_res.Data.Message
	gv_res.ResponseCode = core_res.Status
	if core_res.Data.StatusCode != 200 {
		gv_res.ResponseCode = core_res.Data.StatusCode
	} else {
		gv_res.VaNumber = core_res.Data.VaNo
	}

	return gv_res
}
