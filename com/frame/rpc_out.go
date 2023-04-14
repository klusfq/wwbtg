package frame

// Rpc的输出结构
type RpcOut struct {
	errno  int
	errmsg string
	data   any
}

func NewRpcOut() RpcOut {
	return RpcOut{}
}

func (self *RpcOut) SetData(data any) {
	self.data = data
}

func (self *WxUser) Output(out any) {
	fsout := frame.NewRpcOut()

	fsout.SetData(out)
}
