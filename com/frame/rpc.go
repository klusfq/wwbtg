package frame

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

// Rpc服务
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
