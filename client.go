package main

import (
	"chatroom/client"
	"fmt"
)

func main() {

	var key string
	fmt.Println("Enter your name")
	fmt.Scanf("%s", &key)

	// 將名字傳送給服務器
	userProcess := &client.UserProcess{}
	err := userProcess.Login(key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}



}