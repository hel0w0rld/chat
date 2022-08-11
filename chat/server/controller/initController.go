package controller

import (
	"sync"
	"net"
	"chat/class"
)

type controller struct {
	conn net.Conn
}

// [FILE FUNCTION]
// initialize

// get different controller instance which is bind to user
// in: connection between client and server
// out: controller instance
func GetController(conn net.Conn) *controller {
	return &controller{
		conn : conn,
	}
}

// used for save online users which is bind to the server (global unique)
type onlineUsers struct {
	users map[string]*class.User
	new *class.User
	record []string
}

var online *onlineUsers
var lock sync.Mutex

// get unique online user map which is thread safety in server
// out: unique onlineUsers instance
func GetOnline() *onlineUsers {
	if online == nil {
		lock.Lock()
		lock.Unlock()
		if online == nil {
			online = &onlineUsers{
				users : make(map[string]*class.User, 10),
			}
		}
	}
	return online
}