package view

import (
	"fmt"
	"chat/class"
	"chat/client/controller"
)

// [FILE FUNCTION]
// 1. interact with terminal
// 2. menu UI display 

// choose next process by key
// in: a value used to choose where from menu UI
// out: used to stop outer cycle 
func (this *view) menuRouter(key int) (bool) {
	loop := true
	var err error
	switch key {
		case 1:
			this.onlineUsersUI()
		case 2: 
			err = this.chatUI()
		case 3:
			err = this.broadcastUI()
		case 4:
			err = this.historyUI()
		case 5:
			err = this.exitUI()
		default:
			this.alertUI()
	}
	if err != nil {
		fmt.Println(err)
	}
	return loop
}

// the UI displayed after login success
// out: used to choose next process
func (this *view) menuUI() (key int){
	fmt.Printf("[1].在线列表 [2].私发消息 [3].群发消息 [4].历史记录 [5].退出系统: \n")
	fmt.Scanln(&key)
	return
}

// show the current online users beside me
// out: the number of current online users
func (this *view) onlineUsersUI() (nums int) {
	controller := controller.GetInstance(this.Conn)
	nums = len(controller.Online)
	if nums == 0{
		fmt.Println("无人在线")
		return
	}
	fmt.Println("id\tstate")
	fmt.Println("-------------")

	for id, state := range controller.Online{
		if id == controller.Id{ continue }
		switch state{
			case class.ONLINE:
				fmt.Println(id, "\t在线")
			case class.OFFLINE:
				fmt.Println(id, "\t离线")
		}
	}
	return
}

// choose a user to send private message 
// after show the current online users beside me
func (this *view) chatUI() (err error){
	nums := this.onlineUsersUI()
	if nums == 0{ return }
	
	fmt.Printf("@:")
	var to string
	fmt.Scanln(&to)
	fmt.Printf("[in]:")
	var text string
	fmt.Scanln(&text)

	controller := controller.GetInstance(this.Conn)
	err = controller.Chat(to, text)
	return
}

// send public message to all online users
func (this *view) broadcastUI() (err error){
	fmt.Printf("\r[in]:")
	var text string
	fmt.Scanf("%s\n", &text)

	controller := controller.GetInstance(this.Conn)
	err = controller.Broadcast(text)
	return
}

// request history message from server
func (this *view) historyUI() (err error){
	controller := controller.GetInstance(this.Conn)
	err = controller.History()
	return
}

// request to exit from server
func (this *view) exitUI() (err error){
	controller := controller.GetInstance(this.Conn)
	err = controller.Exit()
	return
}
