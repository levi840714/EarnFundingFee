package binance

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const (
	BaseUrl = "https://fapi.binance.com"
)

func GetTime() {
	uri := BaseUrl + "/fapi/v1/time"
	r, _ := http.NewRequest(http.MethodGet, uri, nil)
	resp, _ := getResponseJson(r)
	var st ServerTime
	_ = json.Unmarshal(resp, &st)
	serverTime := time.Unix(st.ServerTime/1000, 0).Format("2006-01-02 15:04:05")
	fmt.Println(serverTime)
}

func GetFundingRate() (fundingRate []FundingRate) {
	uri := BaseUrl + "/fapi/v1/premiumIndex"
	r, _ := http.NewRequest(http.MethodGet, uri, nil)
	resp, _ := getResponseJson(r)
	_ = json.Unmarshal(resp, &fundingRate)

	return
}

func GetTop10FundingRate() (top10 []FundingRate) {
	fundingRate := GetFundingRate()
	sort.Slice(fundingRate, func(i, j int) bool {
		x, _ := strconv.ParseFloat(fundingRate[i].LastFundingRate, 64)
		y, _ := strconv.ParseFloat(fundingRate[j].LastFundingRate, 64)
		return x > y
	})

	top10 = fundingRate[:10]
	fmt.Printf("%+v", top10)
	return
}

func GetBalance() (usdtBalance decimal.Decimal){
	uri := BaseUrl + "/fapi/v2/balance"
	params := map[string]interface{}{
		"timestamp": milliTimestamp(),
	}
	payload := requestParams(params)
	payload = signature(payload)
	r, _ := http.NewRequest(http.MethodGet, uri+"?"+payload, nil)
	r.Header.Set("X-MBX-APIKEY", viper.GetString("exchange.apiKey"))

	resp, _ := getResponseJson(r)
	var getBalance GetBalanceResp
	_ = json.Unmarshal(resp, &getBalance)
	for _, v := range getBalance {
		if v.Asset == "USDT" {
			usdtBalance, _ = decimal.NewFromString(v.AvailableBalance)
			break
		}
	}
	fmt.Println(usdtBalance)
	return
}

func ChangeLeverage() {

}

func CalculateFundingFee() {

}

func EstimateLiquidationProfit() {

}

func CheckOrderCondition() {

}

func NewOrder() {

}

func ClosePosition() {

}
