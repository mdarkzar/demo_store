package bimport

import (
	"store/internal/bridge"
	"store/tools/datefunctions"
)

// BridgeImports бриджи - это мосты между usecase или интерфейсы для usecase
// которые позволяют безболезненно использовать друг друга
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
