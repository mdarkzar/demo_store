package bimport

import (
	"store/internal/bridge"

	"github.com/golang/mock/gomock"
)

type TestBridgeImports struct {
	ctrl       *gomock.Controller
	TestBridge TestBridge
}

func NewTestBridgeImports(
	ctrl *gomock.Controller,
) *TestBridgeImports {
	return &TestBridgeImports{
		ctrl: ctrl,
		TestBridge: TestBridge{
			Date:         bridge.NewMockDate(ctrl),
			Notification: bridge.NewMockNotification(ctrl),
			Queue:        bridge.NewMockQueue(ctrl),
		},
	}
}

func (t *TestBridgeImports) BridgeImports() *BridgeImports {
	return &BridgeImports{
		Bridge: Bridge{
			Date:         t.TestBridge.Date,
			Notification: t.TestBridge.Notification,
			Queue:        t.TestBridge.Queue,
		},
	}
}
