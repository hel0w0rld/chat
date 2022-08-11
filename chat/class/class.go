package class

import (
	"net"
)

const (
	// Message state
	LOGIN = iota
	REGIST
	EXIT
	HISTORY
	RESULT
	NOTIFY
	CHAT
	BROADCAST
	// User state
	OFFLINE // = iota
	ONLINE 
	BUSY
	
)

type Result struct {
	Code string `json:"code"`
	Describe string `json:"describe"`
	Data string `json:"data"`
}

type Message struct {
	Type int `json:"type"`
	Data string `json:"data"`
}

type User struct {
	Id string `json:"id"`
	State int `json:"state"`
	Conn net.Conn `json:"conn"`
}

type Sms struct {
	From string `json:"from"`
	To string `json:"to"`
	Text string `json:"text"`
}

type Login struct {
	Id string `json:"id"`
	Pwd string `json:"pwd"`
}