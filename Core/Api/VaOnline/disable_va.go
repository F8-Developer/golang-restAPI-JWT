package VaOnline

import (
	"merchant-api-gateway/Core/Structs"
	"merchant-api-gateway/Config"
	"io/ioutil"
	"net/http"
	"encoding/json"
    "time"
	"bytes"
	log "github.com/Sirupsen/logrus"
)

// Register register one new user in db, return a boolean value to make know success or not.
func DisableVa(dv_req Structs.DisableVaRequest) (dv_res Structs.DisableVaResponse) {
	core_secret_key := Config.GoDotEnvVariable("CORE_SECRET_KEY")
	url := Config.GoDotEnvVariable("CORE_URL_DISABLE_VA")

	// send request to core
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	body_req, _ := json.Marshal(map[string]string{
        "merchant_id" : dv_req.MerchantId,
		"merchant_ref_code": dv_req.MerchantRefCode,
		"va_no" : dv_req.VaNumber,
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
	dv_res.ResponseMsg = core_res.Data.Message
	dv_res.ResponseCode = core_res.Status
	if core_res.Data.StatusCode != 200 {
		dv_res.ResponseCode = core_res.Data.StatusCode
	}

	return dv_res
}
