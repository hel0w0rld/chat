package controller
import (
	"net"
)

// [FILE FUNCTION]
// initialize

type controller struct {
	Conn net.Conn
	Id string
	Online map[string]int
}

var instance *controller

// get controller instance
// in: connection where between client and server
// out: unique controller instance
func GetInstance(conn net.Conn) *controller{
	if instance == nil{
		instance = &controller{
			Conn : conn,
		}
	}
	return instance
}