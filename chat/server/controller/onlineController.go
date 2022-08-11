package controller

import (
	"chat/class"
	"errors"
	"time"
	"fmt"
)

// [FILE FUNCTION]
// the CRUD of online users

// create a new online user in users
func (this *onlineUsers) create(){
	this.users[this.new.Id] = this.new
}

// update a new online user in users
// in: user instance which is need to be create or update
func (this *onlineUsers) update(user *class.User){
	this.users[user.Id] = user
}

// delete online user from users
// in: user instance which is need to be delete
func (this *onlineUsers) remove(user *class.User) {
	delete(this.users, user.Id)
}

// check user which named id is online or not
// in: user's id which is used for check
// out: return user instance while user online
func (this *onlineUsers) find(id string) (user *class.User, err error) {
	user, ok := this.users[id]
	if !ok {
		err = errors.New("用户不在线")
	}
	return
}

// get id and state of all online users
// out: a map of user online which is composed of user's id and state
func (this *onlineUsers) getAll() map[string]int{
	onlines := make(map[string]int)
	for _, user := range this.users{
		onlines[user.Id] = user.State
	}
	return onlines
}

// save chat history in server
// in: chat message which need to be save
// out: the number of current chat history
func (this *onlineUsers) save(text string) (nums int){
	nums = len(this.record)
	if nums == 100 {
		this.record = make([]string, 1)
	}
	text = fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 03:04:05"), text)
	this.record = append(this.record, text)
	nums = len(this.record)
	return
}