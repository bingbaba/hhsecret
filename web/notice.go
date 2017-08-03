package web

import (
	"github.com/bingbaba/dingtalk"
	"github.com/bingbaba/hhsecret"
	"github.com/robfig/cron"

	"strings"
)

var (
	UserList                 []string
	DingTalkClient           *dingtalk.DTalkClient
	DefaultDingTalkConfigure DingTalkConfigure
)

func SignCheckLoop(conf DingTalkConfigure, userlist []string) error {
	DefaultDingTalkConfigure = conf
	UserList = userlist

	client, err := dingtalk.NewDTalkClient(conf.CorpID, conf.CorpSecret)
	if err != nil {
		return err
	}
	DingTalkClient = client

	c := cron.New()
	c.AddFunc("0 50 8 * * 1,2,3,4,5", MorningSignCheck)
	c.AddFunc("0 * 23 * * 1,2,3,4,5,6", AfternoonSignCheck)
	c.Start()

	logger.Errorln("init checking...")
	return nil
}

func MorningSignCheck() {
	signCheck(false)
}
func AfternoonSignCheck() {
	logger.Errorln("afternoon checking...")
	signCheck(true)
}

func signCheck(afternoon bool) {
	var userid_notice = make([]string, 0, len(UserList))
	for _, userid := range UserList {
		client, found := GetClientByUser(userid)
		if !found {
			logger.Warnf("userid=%s user not login", userid)
			continue
		}

		lsd, err := client.ListSignPost()
		if err != nil {
			logger.Errorf("userid=%s error=%v get sign list failed", userid, err)
		} else {
			if len(lsd.Signs) == 0 {
				userid_notice = append(userid_notice, userid)
			} else {
				if afternoon {
					max_i := len(lsd.Signs) - 1
					mtime := lsd.Signs[max_i].GetMinuteSecode()
					if strings.Compare(mtime, "17:30") < 0 {
						userid_notice = append(userid_notice, userid)
					}
				}
			}
		}
	}

	notice(userid_notice, afternoon)
}

func newMessage(client *hhsecret.Client, userid string, afternoon bool) *dingtalk.OAMsgContent {
	logger.Infof("userid=%s notice now...", userid)
	richText := "上班"
	if afternoon {
		richText = "下班"
	}

	return &dingtalk.OAMsgContent{
		MsgUrl: "https://m.bingbaba.com/html/sign.html",
		Body: &dingtalk.OAMsgContentBody{
			Title:   "少年，该打卡了...",
			Content: "点击查看详情",
			Rich:    dingtalk.OAMsgContentBodyRich{richText, "打卡了"},
			// Form: []dingtalk.KeyValue{
			// 	dingtalk.KeyValue{"姓名：", client.LoginData.Name},
			// 	dingtalk.KeyValue{"工号：", userid},
			// },
			Image:  "@lADO86ySXc0C1s0FIg",
			Author: "@lazysign",
		},
	}
}

func notice(userids []string, afternoon bool) {
	for _, userid := range userids {
		client, found := GetClientByUser(userid)
		if !found {
			continue
		}

		dtid, found := DefaultDingTalkConfigure.UserID[userid]
		if !found {
			dtid = userid
		}

		target := &dingtalk.MsgTarget{UserIDList: []string{dtid}}
		err := DingTalkClient.SendOAMsg(
			DefaultDingTalkConfigure.AgentID,
			target,
			newMessage(client, userid, afternoon),
		)
		if err != nil {
			logger.Errorf("userid=%s error=%v send oamessage failed!\n", userid, err)
		}
	}
}
