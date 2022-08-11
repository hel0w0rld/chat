package controller
import (
	"chat/client/model"
	"chat/class"
	"errors"
	"encoding/json"
)

// [FILE FUNCTION]
// 1. decode message
// 2. receive related logic processing in menu UI

// wait for message from server
// out: text information from server
func (this *controller) Wait() (describes []string, err error) {
	model := model.GetModel(this.Conn)
	outerJson, err := model.Recive()
	if err != nil { return }
	
	// outer decode
	msg := &class.Message{}
	err = json.Unmarshal(outerJson, &msg)
	if err != nil {
		err = errors.New("[ERR_WAIT_oDECODE] (" + err.Error() + ")")
	}

	switch msg.Type{
		case class.NOTIFY:
			describes, err = this.waitNotify(msg)
		case class.BROADCAST:
			describes, err = this.waitBroadcast(msg)
		case class.HISTORY:
			describes, err = this.waitHistory(msg)
		case class.CHAT:
			describes, err = this.waitChat(msg)
		default:
			err = errors.New("未定义类型")
	}
	return
}

// wait for private message from server
// in: message from server
// out: describe informations which read from message
func (this *controller) waitChat(msg *class.Message) (describes []string, err error) {
	// inner decode
	sms := &class.Sms{}
	err = json.Unmarshal([]byte(msg.Data), sms)
	if err != nil {
		err = errors.New("[ERR_CHAT_iDECODE] (" + err.Error() + ")")
	}

	desc := "[private] " + sms.From + ":" + sms.Text
	describes = append(describes, desc)
	return
}

// wait for history message from server
// in: message from server
// out: describe informations which read from message
func (this *controller) waitHistory(msg *class.Message) (describes []string, err error) {
	// inner deocde
	// describes := make([]string, 10)
	err = json.Unmarshal([]byte(msg.Data), &describes)
	if err != nil {
		err = errors.New("[ERR_HISTORY_iDECODE] (" + err.Error() + ")")
	}
	return
}

// wait for broadcast message from server
// in: message from server
// out: describe informations which read from message
func (this *controller) waitBroadcast(msg *class.Message) (describes []string, err error) {
	// inner decode
	sms := &class.Sms{}
	err = json.Unmarshal([]byte(msg.Data), sms)
	if err != nil {
		err = errors.New("[ERR_BROADCAST_iDECODE] (" + err.Error() + ")")
	}

	desc := "[public] " + sms.From + ":" + sms.Text
	describes = append(describes, desc)
	return
}

// wait for notification message from server
// in: message from server
// out: describe informations which read from message
func (this *controller) waitNotify(msg *class.Message) (describes []string, err error){
	// inner decode
	user := &class.User{}
	// [WARNING] can't recive error which could be create a panic because of connection 
	// information is less in decode process
	json.Unmarshal([]byte(msg.Data), user)
	switch user.State {
		case class.ONLINE:
			this.Online[user.Id] = user.State
			desc := "[上线] " + user.Id 
			describes = append(describes, desc)
		case class.OFFLINE:
			delete(this.Online, user.Id)
			desc := "[下线] " + user.Id
			describes = append(describes, desc)
	}
	return
}