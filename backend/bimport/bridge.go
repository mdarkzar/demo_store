package bimport

import "store/internal/bridge"

type Bridge struct {
	Date         bridge.Date
	Notification bridge.Notification
}

type TestBridge struct {
	Date         *bridge.MockDate
	Notification *bridge.MockNotification
}
