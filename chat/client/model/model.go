package model

import (
	"errors"
)

// [FILE FUNCTION]
// send and receive message between client and server

// recive message from server
// out: data from server by json-encode
func (this *model) Recive() (data []byte, err error){
	data, err = this.send.ReadBytes()
	if err != nil {
		err = errors.New("[ERR_MODEL_RECIVE]:" + err.Error())
	}
	return
}

// send message to server
// in: json-encode data which need to be send
func (this *model) Send(data []byte) (err error) {
	err = this.send.WriteBytes(data)
	if err != nil {
		err = errors.New("[ERR_MODEL_SEND]:" + err.Error())
	}
	return
}

// send & recive (combine with Send() and Recive() )
// in: data which need to be send by json-encode 
// out: json-encode data from server
func (this *model) SendRecive(data []byte) (res []byte, err error){
	err = this.Send(data)
	if err != nil { return }
	
	res, err = this.Recive()
	if err != nil { return }
	return
}