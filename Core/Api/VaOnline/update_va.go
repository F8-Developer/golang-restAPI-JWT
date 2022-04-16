package VaOnline

import (
	"merchant-api-gateway/Core/Structs"
	"merchant-api-gateway/Config"
	"io/ioutil"
	"net/http"
	"encoding/json"
    "time"
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

// Register register one new user in db, return a boolean value to make know success or not.
func UpdateVa(uv_req Structs.UpdateVaRequest) (uv_res Structs.UpdateVaResponse) {
	core_secret_key := Config.GoDotEnvVariable("CORE_SECRET_KEY")
	url := Config.GoDotEnvVariable("CORE_URL_UPDATE_AMOUNT_VA")

	// send request to core
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	amount := fmt.Sprint(uv_req.Amount)
	body_req, _ := json.Marshal(map[string]string{
		"merchant_id" : uv_req.MerchantId,
		"merchant_ref_code": uv_req.MerchantRefCode,
		"va_no" : uv_req.VaNumber,
		"amount" : amount,	  
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
	uv_res.ResponseMsg = core_res.Data.Message
	uv_res.ResponseCode = core_res.Status
	if core_res.Data.StatusCode != 200 {
		uv_res.ResponseCode = core_res.Data.StatusCode
	}

	return uv_res
}
