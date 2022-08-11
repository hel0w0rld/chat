package controller
import (
	"chat/client/model"
	"chat/class"
	"errors"
	"encoding/json"
)

// [FILE FUNCTION]
// 1. encode message
// 2. send related logic processing in menu UI

// send private message to server
// in: target user and content which need to be send 
func (this *controller) Chat(id string, text string) (err error){
	_, ok := this.Online[id]
	if !ok {
		err = errors.New("用户不在线，请重试")
		return
	}
	
	// inner encode
	sms := &class.Sms{
		From: this.Id,
		To: id,
		Text : text,
	}
	innerJson, _ := json.Marshal(sms)
	
	// outer encode
	msg := &class.Message{
		Type : class.CHAT,
		Data : string(innerJson),
	}
	outerJson, _ := json.Marshal(msg)

	// send to server by model
	model := model.GetModel(this.Conn)
	err = model.Send(outerJson)
	if err != nil {
		err = errors.New("[ERR_CONTROL_CHAT]:" + err.Error())
	}
	return
}

// send text that all online users could see to server
// in: content which need to be send
func (this *controller) Broadcast(text string) (err error) {
	// inner encode
	sms := &class.Sms{
		From : this.Id,
		Text : text,
	}
	innerJson, _ := json.Marshal(sms)

	// outer encode
	msg := &class.Message{
		Type : class.BROADCAST,
		Data : string(innerJson),
	}
	outerJson, _ := json.Marshal(msg)

	// send to server by model
	model := model.GetModel(this.Conn)
	err = model.Send(outerJson)
	if err != nil {
		err = errors.New("[ERR_CONTROL_BROADCAST]:" + err.Error())
	}
	return
}

// request history message from server
func (this *controller) History() (err error) {
	// outer encode
	msg := &class.Message{
		Type : class.HISTORY,
	}
	outerJson, _ := json.Marshal(msg)

	// send to server by model
	model := model.GetModel(this.Conn)
	err = model.Send(outerJson)
	if err != nil {
		err = errors.New("[ERR_CONTROL_HISTORY]:" + err.Error())
	}
	return
}

// request exit from server
func (this *controller) Exit() (err error){
	// inner encode
	user := &class.User{
		Id : this.Id,
	}
	innerJson, _ := json.Marshal(user)

	// outer encode
	msg := &class.Message{
		Type : class.EXIT,
		Data : string(innerJson),
	}
	outerJson, _ := json.Marshal(msg)

	// send to server by model
	model := model.GetModel(this.Conn)
	err = model.Send(outerJson)
	if err != nil {
		err = errors.New("[ERR_CONTROL_HISTORY]:" + err.Error())
	}
	return
}