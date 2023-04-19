package config

import (
	"sync"
	"wwbtg/com/frame"
	"wwbtg/user/service"
)

var ServiceMap map[string]frame.BaseService
var once sync.Once

func init() {
	once.Do(func() {
		ServiceMap = map[string]frame.BaseService{
			"WxUser":   new(service.WxUser),
			"AliUser":  new(service.AliUser),
			"Passport": new(service.Passport),
		}
	})
}
