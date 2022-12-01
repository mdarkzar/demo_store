package bimport

import (
	"store/internal/bridge"
	"store/tools/datefunctions"
)

type BridgeImports struct {
	Bridge Bridge
}

func (b *BridgeImports) InitBridge(
	notification bridge.Notification,
) {
	b.Bridge = Bridge{
		Date:         datefunctions.NewDateTool(),
		Notification: notification,
	}
}

func NewEmptyBridge() *BridgeImports {
	return &BridgeImports{}
}
