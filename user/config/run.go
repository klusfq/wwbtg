package config

import (
	"wwbtg/com/frame"
	"wwbtg/user/service"
)

var ServiceMap map[string]frame.BaseService

func init() {
	ServiceMap = map[string]frame.BaseService{
		"WxUser":   new(service.WxUser),
		"AliUser":  new(service.AliUser),
		"Passport": new(service.Passport),
	}
}
