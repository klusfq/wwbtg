package frame

import (
	"net/http"
)

// -- api的action接口
type BaseAction interface {
	Invoke() (string, error)
	Register() (http.Handler, error)
}

// -- service接口
type BaseService interface {
	GetInput(key string, deft any) any
	Output(out any) RpcOut
}
