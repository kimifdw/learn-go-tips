package state

import (
	"testing"
	"time"
)

func TestGetState(t *testing.T) {
	// 初始化状态管理
	stateManagerInstance = &StateManager{}
	stateManagerInstance.initState(300 * time.Millisecond)

	// 模拟发送短信的接口
	sendSms := func() {
		// 发送短信
		GetState().Send(&Context{
			Tel:        "+8613666666666",
			Text:       "3232",
			TemplateID: "TYSHK_01",
		})
	}

	// 模拟用户调用发送短信的接口
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
}
