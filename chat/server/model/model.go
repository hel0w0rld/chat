package model

import(
	"github.com/gomodule/redigo/redis"
	"errors"
)

// [FILE FUNCTION]
// redis operation

// search user in redis or not
// in: the id of user
// out: the password of user who in redis
func (this *model) find(id string) (pwd string, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	pwd, err = redis.String(conn.Do("GET", id))
	if err != nil {
		err = errors.New("[ERR_MODEL_FIND]:用户不存在")
	}
	return
}

// check user in redis or not
// in: the id & password of user
func (this *model) Login(id string, password string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	pwd, err := this.find(id)
	if err != nil { return }
	// login success
	if password != pwd {
		err = errors.New("[ERR_MODEL_LOGIN]:密码错误")
	}
	return 
}

// save user's information in redis
// in: the id & password of user
func (this *model) Regist(id string, pwd string) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.find(id)
	if err != nil {
		conn.Do("SET", id, pwd)
		err = nil
	}else{
		err = errors.New("[ERR_MODEL_REGIST]:用户已存在")
	}
	return
}