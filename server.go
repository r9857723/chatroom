package main

import (
	"chatroom/server"
	"fmt"
	"net"
)

func main() {
	//redis.InitRedisPool("localhost:6379", 16, 0, 300*time.Second)
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("listen port = 8889 fail err = ", err)
		return
	}
	defer listen.Close()

	fmt.Println("listening....")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen fail err = ", err)
		}
		serverProcess := &server.ServerProcess{}
		go serverProcess.ProcessType(conn)
	}
}
