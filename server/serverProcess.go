package server

import (
	"chatroom/common"
	"chatroom/utils"
	"errors"
	"fmt"
	"net"
)

type ServerProcess struct {

}

func (s *ServerProcess) ProcessType(conn net.Conn) (err error) {
	defer conn.Close()

	fmt.Println(conn)
	// 讀取使用者端發送的訊息
	for {
		transfer := &utils.Transfer{
			Conn: conn,
		}
		message, err := transfer.ReadPkg()
		if err != nil {
			return errors.New("transfer.ReadPkg() fail err = " + err.Error())
		}
		// 判斷使用邏輯
		switch message.Type {
			case common.LoginDataType:
				userProcess := &UserProcess{
					Conn: conn,
				}
				_ = userProcess.loginDataProcess(message)
			case common.SmsMessageDataType:
				userProcess := &UserProcess{
					Conn: conn,
				}
				_ = userProcess.SendMessage(message)
			default:
				fmt.Println("無法辨別")

		}
	}
}

