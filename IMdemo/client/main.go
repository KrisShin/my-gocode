package main

import "fmt"

//定义两个全局变量, 一个是用户id, 一个是用户密码
var userId int
var userPwd string

func main() {
	// 接收用户选择
	// 判断是否继续显示菜单

	var key int

	var loop bool = true
	for {
		fmt.Println("------------------欢迎登陆多人聊天室-----------------")
		fmt.Println("\t\t\t 1 登陆聊天室 ")
		fmt.Println("\t\t\t 2 注册用户 ")
		fmt.Println("\t\t\t 3 退出聊天室 ")
		fmt.Println("\t\t\t 请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			fmt.Println("输入有误, 请重新输入")
		}

		//根据用户的输入显示新的提示信息
		if key == 1 {
			// 用户登陆
			fmt.Println("请输入ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			// 把登陆函数写到其他文件
			err := login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败")
			} else {
				fmt.Println("登陆成功")
			}
		} else if key == 2 {
			fmt.Println("用户注册逻辑")
			//用户注册逻辑
		}
	}

}
