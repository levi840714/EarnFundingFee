package binance

type (
	ServerTime struct {
		ServerTime int64 `json:"serverTime"`
	}

	FundingRate struct {
		Symbol               string `json:"symbol"`
		MarkPrice            string `json:"markPrice"`
		IndexPrice           string `json:"indexPrice"`
		EstimatedSettlePrice string `json:"estimatedSettlePrice"`
		LastFundingRate      string `json:"lastFundingRate"`
		NextFundingTime      int64  `json:"nextFundingTime"`
		InterestRate         string `json:"interestRate"`
		Time                 int64  `json:"time"`
	}

	GetBalanceResp []struct {
		AccountAlias       string `json:"accountAlias"`
		Asset              string `json:"asset"`
		Balance            string `json:"balance"`
		CrossWalletBalance string `json:"crossWalletBalance"`
		CrossUnPnl         string `json:"crossUnPnl"`
		AvailableBalance   string `json:"availableBalance"`
		MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	}
)
