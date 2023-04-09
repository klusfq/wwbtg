package frame

import (
	"log"
	"net/http"
)

type Serv struct {
	mod string
	mux *http.ServeMux
}

/**
 * 初始化框架
 *	1. 加载服务配置
 *  2. 加载路由
 */
func (self *Serv) Init(mod, configPath string, apiList map[string]BaseAction) {
	log.Println("Framework init...")
	self.mod = mod
	self.mux = http.NewServeMux()

	// TODO

	for api, ba := range apiList {
		fc, err := ba.Register()
		if err != nil {
			log.Panic(fc)
		}
		self.mux.Handle(api, fc)
	}
}

/**
 * 启动一个服务
 */
func (self *Serv) Run() {
	log.Printf("Framework mod[%s] is start", self.mod)

	realServ := &http.Server{
		Addr:    ":8081",
		Handler: self.mux,
	}

	realServ.ListenAndServe()
}
