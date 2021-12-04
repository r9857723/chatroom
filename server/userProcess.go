package server

import (
	"chatroom/common"
	"chatroom/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	Name string
}


// loginDataProcess 處理登陸邏輯
func (u *UserProcess) loginDataProcess(message common.Message) (err error) {
	// 將接收到的message 反序列化
	var loginData common.LoginData
	err = json.Unmarshal([]byte(message.Data), &loginData)
	if err != nil {
		return errors.New("json.Unmarshal([]byte(message.Data), &a) fail err = " + err.Error())
	}
	fmt.Printf("%s, login success. \n", loginData.UserName)

	// 通知客戶端登入成功
	var loginDataResponse common.LoginDataResponse
	loginDataResponse.Code = 200
	// 回寫使用者名字
	u.Name = loginData.UserName
	// 紀錄線上使用者
	UserMgr.AddOnlineUser(u)
	data, err := json.Marshal(loginDataResponse)
	if err != nil {
		return errors.New("json.Marshal(loginDataResponse) fail	err = " + err.Error())
	}

	var responseMessage common.Message
	responseMessage.Type = common.LoginDataResponseType
	responseMessage.Data = string(data)
	data, err = json.Marshal(responseMessage)
	if err != nil {
		return errors.New("json.Marshal(responseMessage) fail err = " + err.Error())
	}
	transfer := utils.Transfer{
		Conn: u.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		return errors.New("transfer.WritePkg(data) fail err = " + err.Error())
	}

	UserMgr.Show()
	return
}