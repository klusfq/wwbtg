package main

import (
	"encoding/gob"
	"log"
	"net"
	"net/rpc"
	"wwbtg/user/config"
)

func main() {
	rc := rpc.NewServer()

	lisenter, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Panic(err)
	}

	gob.Register(map[string]any{})

	// servList := []string{"WxUser", "AliUser", "Passport"}
	// for _, servName := range servList {
	// }

	for servN, servIns := range config.ServiceMap {
		rc.RegisterName(servN, servIns)
	}

	// rc.RegisterName("WxUser", new(service.WxUser))
	// log.Printf("rpc service[%s] start", "WxUser")

	rc.Accept(lisenter)
}
