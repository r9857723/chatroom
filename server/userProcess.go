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
func (u *UserProcess) SendMessage(message common.Message) (err error) {
	// 將接收到的message 反序列化為可讀資料
	var smsMessageData common.SmsMessageData
	err = json.Unmarshal([]byte(message.Data), &smsMessageData)
	if err != nil {
		return errors.New("json.Unmarshal([]byte(message.Data), &smsMessageData) fail err = " + err.Error())
	}

	data, err := json.Marshal(message)
	if err != nil {
		return errors.New("json.Marshal(message) fail err = " + err.Error())
	}

	fmt.Printf("Name=%s, Context=%s \n", smsMessageData.UserName, smsMessageData.Context)

	// 取得發送者Name
	for name, userProcess := range UserMgr.onlineUser {
		fmt.Println("name=", name)
		fmt.Println("send name = ", smsMessageData.UserName)
		if name == smsMessageData.UserName {
			continue
		}
		transfer := &utils.Transfer{
			Conn: userProcess.Conn,
		}
		err = transfer.WritePkg(data)
		if err != nil {
			return errors.New("transfer.WritePkg(data) fail err = " + err.Error())
		}
	}

	return
}

// loginDataProcess 處理登陸邏輯
func (u *UserProcess) loginDataProcess(message common.Message) (err error) {
	// 將接收到的message 反序列化為可讀資料
	var loginData common.LoginData
	err = json.Unmarshal([]byte(message.Data), &loginData)
	if err != nil {
		return errors.New("json.Unmarshal([]byte(message.Data), &a) fail err = " + err.Error())
	}
	fmt.Printf("%s, login success. \n", loginData.UserName)

	// 通知使用者端登入成功
	var loginDataResponse common.LoginDataResponse
	loginDataResponse.Code = 200
	// 回寫使用者名字
	u.Name = loginData.UserName
	// 紀錄線上使用者
	UserMgr.AddOnlineUser(u)

	transfer := utils.Transfer{
		Conn: u.Conn,
	}
	// 將訊息序列化回覆給使用者
	var responseMessage common.Message
	responseMessage.Type = common.LoginDataResponseType
	data, err := json.Marshal(loginDataResponse)
	if err != nil {
		return errors.New("json.Marshal(loginDataResponse) fail	err = " + err.Error())
	}
	responseMessage.Data = string(data)
	data, err = json.Marshal(responseMessage)
	if err != nil {
		return errors.New("json.Marshal(responseMessage) fail err = " + err.Error())
	}
	err = transfer.WritePkg(data)
	if err != nil {
		return errors.New("transfer.WritePkg(data) fail err = " + err.Error())
	}

	// 顯示在線使用者
	UserMgr.Show()
	return
}