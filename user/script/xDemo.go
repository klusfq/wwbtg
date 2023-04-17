package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"wwbtg/com/frame"
)

func main() {
	// frame.LoadFrameConf("hello")

	log.Println(os.Getwd())
}

func getAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

func tServ() {
	client, err := rpc.Dial("tcp", "localhost:8088")

	if err != nil {
		log.Fatal("dailing: ", err)
	}
	gob.Register(frame.RpcIn{})
	gob.Register(map[string]any{})

	var reply frame.RpcOut
	input := frame.RpcIn{}

	if err := client.Call("WxUser.LoginWxUserService", input, &reply); err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}