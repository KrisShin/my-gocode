package main

import (
	"fmt"
	"net"
	"json"
)

func login(userID int, userPwd string)(err error){
	// 制定协议
	// fmt.Printf("userId = %d userPwd = %s\n",userID,userPwd)
	// return 

	// 1. 连接服务器
	conn,err:=net.Dial("tcp", "localhost:8889")
	if err!=nil{
		fmt.Println("net,Dial err=", err)
		return
	}

	// 2. 准备通过conn发送消息给服务器
	var msg message.Message
	msg.Type = message.LoginMsgType

	// 3. 创建一个loginmsg 结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	//4. 将loginMsg 序列化
	data, err:= json.Marshal(loginMsg)
	if err!=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	msg.Data = string(data)
}
