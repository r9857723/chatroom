package client

import (
	"bufio"
	"chatroom/common"
	"chatroom/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

var CurUser common.CurUser

type UserProcess struct {

}

func (u *UserProcess) readMessage(conn net.Conn) {
	for {
		transfer := utils.Transfer{
			Conn: conn,
		}
		message, err := transfer.ReadPkg()
		if err != nil {
			return
		}
		var smsMessageData common.SmsMessageData
		err = json.Unmarshal([]byte(message.Data), &smsMessageData)
		if err != nil {
			return
		}
		fmt.Printf("%s \t: %s \n", smsMessageData.UserName, smsMessageData.Context)
	}
}

func (u *UserProcess) SendMessage(context string) (err error) {
	var message common.Message
	message.Type = common.SmsMessageDataType
	var smsMessage common.SmsMessageData
	smsMessage.Context = context
	smsMessage.UserName = CurUser.UserName

	data, err := json.Marshal(smsMessage)
	if err != nil {
		return errors.New("json.Marshal(smsMessage) fail err = " + err.Error())
	}
	message.Data = string(data)
	data, err = json.Marshal(message)
	if err != nil {
		return errors.New("json.Marshal(message) fail err = " + err.Error())
	}
	transfer := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		return errors.New("transfer.WritePkg(data) fail err = " + err.Error())
	}
	return
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
		CurUser.UserName = userName
		CurUser.Conn = conn

		go u.readMessage(conn)
		reader := bufio.NewReader(os.Stdin)
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				return errors.New("reader.ReadString('\\n') fail err = " + err.Error())
			}
			str = strings.Replace(str, "\n", "", -1)
			err = u.SendMessage(str)
			if err != nil {
				return errors.New("u.SendMessage(key) fail err = " + err.Error())
			}
		}
	} else {
		fmt.Println("login fail")
	}
	return
}