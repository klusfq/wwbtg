package frame

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/rpc"
)

/**
 * 实现基础的service服务
 */

// Rpc的输入结构
type RpcIn map[string]any

func NewRpcIn() RpcIn {
	return RpcIn{}
}

func (self RpcIn) Get(key string, def any) any {
	val, err := self[key]

	if !err {
		return def
	}

	return val
}

func RpcInBuilder(inner map[string]any) RpcIn {
	return RpcIn(inner)
}

// Rpc的输出结构
type RpcOut struct {
	Errno  int
	Errmsg string
	Data   any
}

func NewRpcOut() RpcOut {
	return RpcOut{}
}

func (self *RpcOut) SetData(data any) {
	self.Data = data
}

// Rpc - 响应端
type DefaultRpcService struct {
	RpcIn
}

func (self *DefaultRpcService) GetInput(key string, def any) any {
	return self.Get(key, def)
}

func (self *DefaultRpcService) Output(out any) RpcOut {
	fsout := NewRpcOut()

	fsout.SetData(out)

	return fsout
}

// Rpc - 请求端
type RpcClient struct {
	client *rpc.Client
}

func NewRpcClient() RpcClient {
	return RpcClient{}
}

func (self *RpcClient) init() {
	var err error
	self.client, err = rpc.Dial("tcp", "localhost:8088")
	gob.Register(map[string]any{})

	if err != nil {
		log.Fatal("dailing: ", err)
	}
}

func (self *RpcClient) Call(module, servApi string, input RpcIn) (*RpcOut, error) {
	self.init()
	var reply RpcOut

	// -- 格式化
	remote := fmt.Sprintf("%s.%s", module, servApi)

	// -- 发送rpc请求
	if err := self.client.Call(remote, input, &reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
