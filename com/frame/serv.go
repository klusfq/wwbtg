package frame

import (
	"log"
	"math/rand"
	"net/http"
	"time"
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
		self.mux.Handle(api, wrapper(fc))
	}
}

// 并发模型
func wrapper(hfc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		oid := rand.Int31n(5)
		// -- 利用context同步子任务同步
		// ctx, cancelFn := context.WithCancel(context.Background())
		// go func() {

		log.Printf("--- local task will cost: %d", oid)
		time.Sleep(time.Second * time.Duration(oid))
		hfc.ServeHTTP(w, r)

		// }()
	})
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
