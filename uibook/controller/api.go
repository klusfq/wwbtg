package controller

import (
	"sync"
	"wwbtg/com/frame"
	"wwbtg/uibook/action"
)

var ApiList map[string]frame.BaseAction
var once sync.Once

func init() {
	once.Do(func() {
		ApiList = map[string]frame.BaseAction{
			"/getbooklist": new(action.GetBookList),
		}
	})
}
