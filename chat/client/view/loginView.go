package view

import (
	"fmt"
	"chat/client/controller"
)

// [FILE FUNCTION]
// 1. interact with terminal 
// 2. login UI display 

// choose next process by key
// in: a value used to choose where from index UI
// out: used to stop the outer cycle 
func (this *view) indexRouter(key int) (bool) {
	loop := true
	switch key{
		case 1:
			// stop outer cycle while login success
			loop = this.loginUI()
		case 2:
			this.registUI()
		default:
			this.alertUI()
	}
	return loop
}


// index UI
// out: used to choose next process
func (this *view) indexUI() (key int) {
	fmt.Printf("[1].登录 [2].注册 [3].退出: ")
	fmt.Scanln(&key)
	return
}

// the same UI in login & regist
// in: used to display login UI or regist UI
// out: id & password
func inputUI(hint string) (id string, pwd string){
	fmt.Printf("[%s]\n",hint)
	fmt.Printf("\r用户:")
	fmt.Scanln(&id)
	fmt.Print("\r密码:")
	fmt.Scanln(&pwd)
	return
}

// login UI
// out: used to stop the outer cycle 
func (this *view) loginUI() (loop bool) {
	id, pwd := inputUI("登录")

	controller := controller.GetInstance(this.Conn)
	ok, describe :=controller.Login(id, pwd)
	if ok {
		loop = false // exit cycle
	}else{
		fmt.Printf("[登录失败]: %s, 请重试\n", describe)
		loop = true // continue cycle
	}
	return
}

// regist UI, continue cycle whether regist success or not
func (this *view) registUI(){
	id, pwd := inputUI("注册")

	controller := controller.GetInstance(this.Conn)
	ok, describe := controller.Regist(id, pwd)
	if ok{
		fmt.Println("[注册成功]")
	}else{
		fmt.Println("[注册失败]: ", describe)
	}
}

// default UI
func (this *view) alertUI() {
	fmt.Println("[?]: 请重新输入")
}
