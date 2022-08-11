package view

import (
	"chat/server/controller"
	"chat/utils"
	"net"
	"io"
	"log"
)

// [FILE FUNCTION]
// main logic control in view part
// PS: although view part in server is not necessary, 
// in order to peeling write & read about client, we keep it. 

// main logic
func Start() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8888")
	defer listener.Close()
	log.Println("[START]")
	for {
		log.Println("[WAIT_NEXT]...")
		conn, err := listener.Accept()
		
		if err != nil {
			log.Println("[ERR_CON]:", err)
		}
		log.Println("[ESTABLISH]:", conn.RemoteAddr())
		go Run(conn)
	}
}

// process different data from client by goroutine
// in: connection between client and server
func Run(conn net.Conn){
	controller := controller.GetController(conn)
	for {
		// "read" in utils.Send is used for once beacuse the connection is different from "write"
		dataJson, err := (&utils.Send{Conn: conn}).ReadBytes()
		if err != nil{
			// client normal exit
			if err == io.EOF{
				log.Printf("[EXIT] %s\n", conn.RemoteAddr())
				return
			}
			log.Println("[ERR_ROUTER]:", err.Error())
		}
		
		// controller router
		datas, err := controller.Router(dataJson)
		if err != nil {
			log.Println(err)
		}
		
		// send data to client
		for conn, dataJson := range datas{
			err = (&utils.Send{Conn: conn}).WriteBytes(dataJson)
			if err != nil {
				log.Println(err)
			}
		}
	}
}