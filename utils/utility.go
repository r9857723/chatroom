package utils

import (
	"chatroom/common"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (T *Transfer) ReadPkg() (mes common.Message, err error) {
	_, err = T.Conn.Read(T.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}

	// 根據buf[:4] 轉成一個uint32
	pkgLen := binary.BigEndian.Uint32(T.Buf[:4])

	// 根據 pkgLen 讀取消息內容
	n, err := T.Conn.Read(T.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	err = json.Unmarshal(T.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("rjson.Unmarshal error")
		return
	}
	return
}

func (T *Transfer) WritePkg(data []byte) (err error) {
	// 先發送長度給對方
	pkgLen := uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(T.Buf[0:4], pkgLen)

	n, err := T.Conn.Write(T.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail err=", err)
		return
	}
	// fmt.Printf("Client send message length success (len=%d) content=%s \n", len(data), string(data))

	// 發送數據
	n, err = T.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail err=", err)
		return
	}
	return
}

/*

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
*/
