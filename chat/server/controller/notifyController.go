package controller

import (
	"net"
	"chat/class"
	"encoding/json"
	"errors"
)

// [FILE FUNCTION]
// 1. sms related operations
// 2. inner encode & decode of data

// notify other online users, a new user online
// in: a new online user
// out: multi key-values of online users
func (this *controller) Notify(user *class.User) (datas map[net.Conn]*class.Message){
	online := GetOnline()
	datas = make(map[net.Conn]*class.Message, len(online.users))
	for id, other := range online.users {
		if id == user.Id { continue }
		// inner encode
		innerJson, _ := json.Marshal(user)
		// outer
		msg := &class.Message{
			Type : class.NOTIFY,
			Data : string(innerJson),
		}
		datas[other.Conn] = msg
	}
	return
}

// forward message from original user to target user
// in: private chat related message information after outer decode from original user 
// out: one key-value of target user
func (this *controller) Chat(msg *class.Message) (datas map[net.Conn]*class.Message, err error){
	datas = make(map[net.Conn]*class.Message, len(online.users))
	// inner decode
	sms := class.Sms{}
	err = json.Unmarshal([]byte(msg.Data), &sms)
	if err != nil {
		err = errors.New("[ERR_CHAT_iDECODE] (" + err.Error() + ")")
		return nil, err
	}	
	
	online := GetOnline()
	user, err := online.find(sms.To)
	if err != nil {return nil, err}

	datas[user.Conn] = msg
	return
}

// forward message without inner decode
// in: public chat related message information after outer decode from client 
// out: multi key-values of online user
func (this *controller) Broadcast(msg *class.Message) (datas map[net.Conn]*class.Message, err error){
	online := GetOnline()
	this.Record(msg)
	datas = make(map[net.Conn]*class.Message, len(online.users))
	for _, other := range online.users{
		datas[other.Conn] = msg
	}
	return
}

// get public chat history since server start
// out: multi key-values of chat history
func (this *controller) History() (datas map[net.Conn]*class.Message){
	online := GetOnline()
	datas = make(map[net.Conn]*class.Message, len(online.users))
	history := online.record
	
	// inner encode
	innerJson, _ := json.Marshal(history)
	// outer
	msg := &class.Message{
		Type : class.HISTORY,
		Data : string(innerJson),
	}
	datas[this.conn] = msg
	return
}

// save public chat history since server start
// in: chat histroy related message information after outer decode from client
func (this *controller) Record(msg *class.Message) {
	// inner decode
	sms := &class.Sms{}
	json.Unmarshal([]byte(msg.Data), sms)
	text := sms.From + ":" + sms.Text

	online := GetOnline()
	online.save(text)
}