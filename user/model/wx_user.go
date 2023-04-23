package model

import (
	"fmt"
	"wwbtg/com/frame"
)

type WxUser struct {
	frame.DefaultModel

	UserId   int    `json:"user_id"`
	RealName string `json:"real_name"`
	WxUniqId int    `json:"wx_uniq_id"`
	WxNick   string `json:"wx_nick"`
	WxGender uint8  `json:"wx_gender"`
}

func NewWxUser() *WxUser {
	out := new(WxUser)
	out.InitDB("wx_user")
	return out
}

func GetWxUserListByName(name string) {
	daoWxUser := NewWxUser()

	fields := []string{
		"user_id",
		"wx_nick",
		"real_name",
	}
	conds := []string{
		fmt.Sprintf("`real_name` = '%s'", name),
		"`wx_gender` = 0",
	}

	var outList []map[string]any
	rows := daoWxUser.Select(fields, conds)
	for rows.Next() {
		rx := new(WxUser)

		rows.Scan(&rx.UserId, &rx.WxNick, &rx.RealName)

		outList = append(outList, map[string]any{
			"userId":   rx.UserId,
			"wxNick":   rx.WxNick,
			"realName": rx.RealName,
		})
	}

	fmt.Println(outList)
}
