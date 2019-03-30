package main

import (
	"fmt"

	"github.com/fictionbase/fictionbase"
	"github.com/nlopes/slack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	cw := fictionbase.NewCw()
	token := viper.GetString("slack.token")
	channel := viper.GetString("slack.channel")
	if token == "" || channel == "" {
		panic("please set Token and channel")
	}
	// @TODO get MetricsParam From config
	val := cw.GetCloudWatch("hoge", "huga", "foo")
	fmt.Println(val.GoString())
	api := slack.New(token)
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Event",
				Value: "TypeKey",
			},
			slack.AttachmentField{
				Title: "Value",
				Value: val.GoString(),
			},
		},
	}
	params.Username = "FictionBase"
	_, _, err := api.PostMessage(channel, slack.MsgOptionText("", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Error(err.Error())
		return
	}
}
