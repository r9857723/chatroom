package server

import "fmt"

var UserMgr *UserMgt

type UserMgt struct {
	onlineUser map[string]*UserProcess
}

// 服務器啟動時初始化
func init() {
	UserMgr = &UserMgt{
		onlineUser: make(map[string]*UserProcess, 1024),
	}
}

// AddOnlineUser 新增使用者
func (u *UserMgt) AddOnlineUser(userProcess *UserProcess) {
	UserMgr.onlineUser[userProcess.Name] = userProcess
}

// DeleteOnlineUser 刪除使用者
func (u *UserMgt) DeleteOnlineUser(Name string) {
	delete(UserMgr.onlineUser, Name)
}

func (u *UserMgt) Show() {
	fmt.Println(UserMgr)
}