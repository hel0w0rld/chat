package main

import (
	"github.com/gomodule/redigo/redis"
	"chat/server/view"
	"chat/server/controller"
	"chat/server/model"
)

func init() {
	initPool("127.0.0.1:6379")
	controller.GetOnline()
}

func initPool(address string){
	model.Pool = &redis.Pool{
		MaxIdle: 8,
		MaxActive: 0, 
		IdleTimeout: 100,
		Dial : func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	
}

func main() {
	view.Start()
}