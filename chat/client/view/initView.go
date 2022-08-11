package view

import (
	"net"
)

// [FILE FUNCTION]
// initialize

type view struct {
	Conn net.Conn
}

var instance *view

// get view instance
// in: connection where between client and server
// out: global unique view instance
func GetInstance(conn net.Conn) *view{
	if instance == nil{
		instance = &view{
			Conn : conn,
		}
	}
	return instance
}