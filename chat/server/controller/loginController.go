package controller

import (
	"encoding/json"
	"errors"
	"chat/class"
	"chat/server/model"
	"net"
)

// [FILE FUNCTION]
// 1. login related operations
// 2. inner encode & decode of data

// write regist information which get from client into redis
// in: user regist related message information after outer decode from client 
// out: one key-value of user regist
func (this *controller) Regist(msg *class.Message) (map[net.Conn]*class.Message, error){
	datas := make(map[net.Conn]*class.Message, 1)
	
	// inner decode
	regist := class.Login{}
	err := json.Unmarshal([]byte(msg.Data), &regist)
	if err != nil {
		err = errors.New("[ERR_REGIST_iDECODE] (" + err.Error() + ")")
		return nil, err
	}

	// write to redis
	res := class.Result{}
	model := model.GetModel()
	err = model.Regist(regist.Id, regist.Pwd)
	if err != nil {
		res.Code = "500"
		res.Describe = "用户已存在"
	} else {
		res.Code = "200"
		res.Describe = "注册成功"
	}

	// inner encode
	innerJson, _ := json.Marshal(res)
	// outer
	msg = &class.Message{
		Type: class.RESULT,
		Data: string(innerJson),
	}

	datas[this.conn] = msg
	return datas, err
}

// check login information which get from client in redis
// in: user login related message information after outer decode from client 
// out: one key-value of user login
func (this *controller) Login(msg *class.Message) (map[net.Conn]*class.Message, error){
	datas := make(map[net.Conn]*class.Message, 1)
	// inner decode
	login := class.Login{}
	err := json.Unmarshal([]byte(msg.Data), &login)
	if err != nil {
		err = errors.New("[ERR_LOGIN_iDECODE] (" + err.Error() + ")")
		return nil, err
	}

	res := class.Result{}
	user, _ := online.find(login.Id)
	if user == nil {
		model := model.GetModel()
		err = model.Login(login.Id, login.Pwd)
		if err != nil {// login fail
			res.Code = "500"
			res.Describe = "用户不存在或密码错误"
		} else {// login success
			online.new = &class.User{
				Id : login.Id,
				State : class.ONLINE,
				Conn : this.conn,
			}

			res.Code = "200"
			res.Describe = "登录成功"
			onlineJson, _ := json.Marshal(online.getAll())
			res.Data = string(onlineJson)
		}
	}else{
		res.Code = "500"
		res.Describe = "用户已登录"
	}

	// inner encode
	resJson, _ := json.Marshal(res)
	// outer
	msg = &class.Message{
		Type: class.RESULT,
		Data: string(resJson),
	}

	datas[this.conn] = msg

	return datas, err
}

// user normal exit
// in: user exit related message information after outer decode from client 
// out: user instance
func (this *controller) Exit(msg *class.Message) *class.User{
	// inner decode
	user := &class.User{}
	json.Unmarshal([]byte(msg.Data), &user)
	user.State = class.OFFLINE

	online := GetOnline()
	online.remove(user)

	return user
}
