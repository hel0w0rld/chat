package main

import (
	"chat/client/view"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	defer conn.Close()

	if err != nil {
		fmt.Println("[ERR_MAIN_INIT]: (" + err.Error() + ")")
		return
	}
	
	view := view.GetInstance(conn)
	view.Start()
}