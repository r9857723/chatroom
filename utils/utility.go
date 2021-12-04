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