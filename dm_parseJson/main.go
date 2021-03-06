package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func main() {
	// var s Serverslice
	// str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	// json.Unmarshal([]byte(str), &s)
	// fmt.Println(s)

	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "sichuang", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "shanghai", ServerIP: "127.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
