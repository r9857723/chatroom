package common

const (
	LoginDataType = "LoginData"
	LoginDataResponseType = "LoginDataResponse"
	SmsMessageDataType = "SmsMessageData"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// LoginData 登入 使用者 -> 服務器
type LoginData struct {
	User
}

// LoginDataResponse 登入 服務器 -> 使用者
type LoginDataResponse struct {
	Code int `json:"code"`
	Err string `json:"err"`
}

// SmsMessageData 發送訊息
type SmsMessageData struct {
	Context string `json:"context"`
	User
}