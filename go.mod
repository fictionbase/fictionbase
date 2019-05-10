module github.com/fictionbase/fictionbase

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/aws/aws-sdk-go v1.19.27
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/mitchellh/go-ps v0.0.0-20170309133038-4fdf99ab2936
	github.com/nlopes/slack v0.5.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.3.0
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190509222800-a4d6f7feada5 // indirect
)

replace (
	github.com/fictionbase/agent => ../agent
	github.com/fictionbase/fictionbase => ../fictionbase
	github.com/fictionbase/monitor => ../monitor
	github.com/fictionbase/router => ../router
)
