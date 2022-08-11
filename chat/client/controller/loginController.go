package controller
import (
	"chat/client/model"
	"chat/class"
	"errors"
	"encoding/json"
)

// [FILE FUNCTION]
// 1. encode & decode message
// 2. login related operations

// user login
// in: id & password
// login success or not
func (this *controller) Login(id string, pwd string) (ok bool, err error) {
	// inner encode
	user := &class.Login{
		Id : id,
		Pwd : pwd,
	}
	innerJson, _ := json.Marshal(user)

	// outer encode
	msg := &class.Message{
		Type : class.LOGIN,
		Data : string(innerJson),
	}
	outerJson, _ := json.Marshal(msg)

	// interact with server by model
	model := model.GetModel(this.Conn)
	outerJson, err = model.SendRecive(outerJson)
	if err != nil {
		err = errors.New("[ERR_CONTROL_LOGIN]:" + err.Error())
		return
	}
	
	// outer decode
	outer := &class.Message{}
	err = json.Unmarshal(outerJson, &outer)
	if err != nil {
		err = errors.New("[ERR_LOGIN_oDECODE] (" + err.Error() + ")")
	}

	// inner decode
	res := &class.Result{}
	err = json.Unmarshal([]byte(outer.Data), res)
	if err != nil {
		err = errors.New("[ERR_LOGIN_iDECODE] (" + err.Error() + ")")
	}

	// login success
	if res.Code == "200"{
		ok = true
		// update information of current user's information
		this.Id = id
		err = json.Unmarshal([]byte(res.Data), &this.Online)
		if err != nil {
			err = errors.New("[ERR_LOGIN_iDECODE] (" + err.Error() + ")")
		}
	}else{
		err = errors.New(res.Describe)
	}
	return
}

// user regist
// in: id & password
// regist success or not
func (this *controller) Regist(id string, pwd string) (ok bool, err error) {
	// inner encode
	user := &class.Login{
		Id : id,
		Pwd : pwd,
	}
	innerJson, _ := json.Marshal(user)

	// outer encode
	msg := &class.Message{
		Type : class.REGIST,
		Data : string(innerJson),
	}
	outerJson, _ := json.Marshal(msg)

	// interact with server by model
	model := model.GetModel(this.Conn)
	outerJson, err = model.SendRecive(outerJson)
	if err != nil {
		err = errors.New("[ERR_CONTROL_REGIST]:" + err.Error())
		return
	}

	// outer decode
	outer := &class.Message{}
	err = json.Unmarshal(outerJson, &outer)
	if err != nil {
		err = errors.New("[ERR_REGIST_oDECODE] (" + err.Error() + ")")
	}

	// inner decode
	res := &class.Result{}
	err = json.Unmarshal([]byte(outer.Data), res)
	if err != nil {
		err = errors.New("[ERR_REGIST_iDECODE] (" + err.Error() + ")")
	}

	// regist success
	if res.Code == "200"{
		ok = true
	}else{
		err = errors.New(res.Describe)
	}
	return
}