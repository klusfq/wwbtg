package service

import (
	"fmt"
	"wwbtg/com/frame"
)

type WxUser struct {
	frame.DefaultRpcService
}

func (self *WxUser) LoginWxUserService(req frame.RpcIn, resp *frame.RpcOut) error {
	out := map[string]any{
		"userInfo": map[string]any{
			"name": "fq",
			"age":  12,
		},
	}

	fmt.Println(out)

	*resp = self.Output(out)

	return nil
}

// -- 实现BaseService
