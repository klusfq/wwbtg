package controller

import (
	"wwbtg/com/frame"
	"wwbtg/uibook/action"
)

var ApiList map[string]frame.BaseAction

func init() {
	ApiList = map[string]frame.BaseAction{
		"/getbooklist": &action.GetBookList{},
	}
}
