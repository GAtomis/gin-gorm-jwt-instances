package model

type GetTokenReq struct {
	Userid    string `json:"userid"`
	Appid     string `json:"appid"`
	AppSecert string `json:"appsecrt"`
}

type SendGiftReq struct {
	Timestamp    int64  `json:"timestamp"`
	Expiretimems int64  `json:"expiretimems"`
	Userid       string `json:"userid"`
	Roomid       string `json:"roomid"`

	Appid  string `json:"appid"`
	Gameid string `json:"gameid"`

	GiftInfo GiftMode `json:"data"`
}

type Resp struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Ts        int         `json:"ts"`
	RequestId string      `json:"requestId"`
	Data      interface{} `json:"data"`
}

type GiftMode struct {
	To      string      `json:"to"`
	Giftobj GiftPayload `json:"payload"`
}

type GiftPayload struct {
	GiftCost int32 `json:"giftCost"`
	Count    int32 `json:"count"`
}
