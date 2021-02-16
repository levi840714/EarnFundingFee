package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func milliTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func getResponseJson(req *http.Request) (jsonByte []byte, err error) {
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	jsonByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func requestParams(params map[string]interface{}) string {
	payload := url.Values{}
	for k, v := range params {
		v := fmt.Sprintf("%v", v)
		payload.Set(k, v)
	}

	return payload.Encode()
}

func signature(payload string) string {
	h := hmac.New(sha256.New, []byte(viper.GetString("exchange.secretKey")))
	h.Write([]byte(payload))
	sign := hex.EncodeToString(h.Sum(nil))
	payload = payload + "&signature=" + sign

	return payload
}
