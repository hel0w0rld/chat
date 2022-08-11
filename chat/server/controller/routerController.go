package controller

import (
	"net"
	"errors"
	"chat/class"
	"encoding/json"
)

// [FILE FUNCTION]
// 1. outer encode & decode of data
// 2. function selection in server

// decode information from client and function selection in server
// in: data from client
// out: multi key-value in this function
func (this *controller) Router(data []byte) (datasJson map[net.Conn][]byte, err error) {
	
	// ouoter decode
	msg := &class.Message{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		err = errors.New("[ERR_ROUTER_oDECODE] (" + err.Error() + ")")
		return
	}
	
	// router
	datas := make(map[net.Conn]*class.Message, 1)
	switch msg.Type{
		case class.LOGIN:
			datas, err = this.Login(msg)
			if err != nil {break}
			//update online information
			online := GetOnline()
			online.create()
			others := this.Notify(online.new)
			// combine datas with others
			for key, value := range others {
				datas[key] = value
			}
		case class.REGIST:
			datas, err = this.Regist(msg)
		case class.BROADCAST:
			datas, err = this.Broadcast(msg)
		case class.HISTORY:
			datas = this.History()
		case class.CHAT:
			datas, err = this.Chat(msg)
		case class.EXIT:
			exit := this.Exit(msg)
			datas = this.Notify(exit)
		default:
			err = errors.New("未知类型")
	}

	// outer encode
	datasJson = make(map[net.Conn][]byte, 1)
	for key, value := range datas {
		dataJson, _ := json.Marshal(value)
		datasJson[key] = dataJson
	}
	return
}
