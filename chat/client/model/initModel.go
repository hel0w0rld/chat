package model

import (
	"chat/utils"
	"net"
)

// [FILE FUNCTION]
// initialize

type model struct {
	send *utils.Send
}

var modelInstance *model

// get model instance
// in: connection where between client and server
// out: global unique model instance
func GetModel(conn net.Conn) *model {
	if modelInstance == nil {
		modelInstance = &model{
			send : &utils.Send{
				Conn : conn,
			},
		}
	}
	return modelInstance
}