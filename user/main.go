package main

import (
	"encoding/gob"
	"log"
	"net"
	"net/rpc"
	"os"
	"wwbtg/user/service"
)

func main() {
	rc := rpc.NewServer()

	lisenter, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Panic(err)
	}

	gob.Register(map[string]any{})
	rc.RegisterName("WxUser", new(service.WxUser))

	// log.Printf("rpc service[%s] start", "WxUser")
	rc.Accept(lisenter)
}

func LoadAppConf(fnm string) {
	po := os.Getwd()
	fname := po + "/" + fnm + ".json"

}
