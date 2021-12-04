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
	// 讀取客戶端發送的訊息
	for {
		transfer := &utils.Transfer{
			Conn: conn,
		}
		message, err := transfer.ReadPkg()
		if err != nil {
			return errors.New("transfer.ReadPkg() fail err = " + err.Error())
		}
		fmt.Println(message)
		// 判斷使用哪個邏輯
		switch message.Type {
			case common.LoginDataType:
				userProcess := &UserProcess{
					Conn: conn,
				}
				_ = userProcess.loginDataProcess(message)
		}
	}
}

