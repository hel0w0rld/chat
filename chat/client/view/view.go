package view

import (
	"fmt"
	"chat/client/controller"
)

// [FILE FUNCTION]
// main logic control in view part

// main logic
func (this *view) Start(){
	// index UI
	loop := true
	for loop {
		key := this.indexUI()
		// exit
		if key == 3 {return}
		loop = this.indexRouter(key)
	}

	// wait message from server
	controller := controller.GetInstance(this.Conn)
	go func(){
		for {
			describes, err := controller.Wait()
			if err != nil{
				fmt.Println(err)
			}else{
				for _, describe := range describes{
					fmt.Println(describe)
				}
			}
		}
	}()
	
	// menu UI
	loop = true 
	for loop{
		key := this.menuUI()
		loop = this.menuRouter(key)
		// exit
		if key == 5 {return}
	}
}