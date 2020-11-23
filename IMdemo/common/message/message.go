package message

// 消息类型常量
const (
	LoginMsgType    = "LoginMsg"
	LoginResMsgType = "LoginResMsg"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息的内容
}

// 定义两个消息
type LoginMsg struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  // 用户米阿曼
	UserName string `json:"userName"` // 用户名
}

type LoginResMsg struct {
	Code  int    `json:"code"`  // 返回状态码 500-未注册 200-成功
	Error string `json:"error"` // 返回的错误信息
}
