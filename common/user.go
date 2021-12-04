package common

import "net"

type User struct {
	UserName string `json="username"`
}

type CurUser struct {
	Conn net.Conn
	User
}