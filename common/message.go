package common

const (
	LoginDataType = "LoginData"
	LoginDataResponseType = "LoginDataResponse"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginData struct {
	User
}

type LoginDataResponse struct {
	Code int `json:"code"`
	Err string `json:"err"`
}