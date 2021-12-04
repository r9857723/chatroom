package client

import (
	"chatroom/common"
	"chatroom/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {

}

func (u *UserProcess) Login(userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return errors.New("connect to server fail err = " + err.Error())
	}
	defer conn.Close()

	transfer := &utils.Transfer{
		Conn: conn,
	}
	// 1.通過 conn 發送名字給服務器
	var message common.Message
	message.Type = common.LoginDataType
	var loginData common.LoginData
	loginData.UserName = userName

	data, err := json.Marshal(loginData)
	if err != nil {
		return errors.New("json.Marshal(loginData) fail err = " + err.Error())
	}
	message.Data = string(data)
	data, err = json.Marshal(message)
	if err != nil {
		return errors.New("json.Marshal(message) fail err = " + err.Error())
	}
	err = transfer.WritePkg(data)
	if err != nil {
		return errors.New("transfer.WritePkg(data) fail err = " + err.Error())
	}

	// 2.接收服務器訊息
	result, err := transfer.ReadPkg()
	if err != nil {
		return errors.New("transfer.ReadPkg() fail err = " + err.Error())
	}
	var loginDatsResponse common.LoginDataResponse
	err = json.Unmarshal([]byte(result.Data), &loginDatsResponse)
	if err != nil {
		return errors.New("json.Unmarshal([]byte(result.Data), &loginDatsResponse) fail err = " + err.Error())
	}

	if loginDatsResponse.Code == 200 {
		fmt.Println("login success")
	} else {
		fmt.Println("login fail")
	}
	return
}