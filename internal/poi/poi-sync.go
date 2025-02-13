package poi

import (
	"time"
)

// PoiSync struct
type PoiSync struct {
	isShutdown      bool
	latestSyncedPoi *PoiBlock
	lastFlushTs     time.Time
	isSyncing       bool
	syncedPoiQueue  []PoiBlock // Using a slice for queue functionality
}

// Method to implement shutdown logic
func (p *PoiSync) OnApplicationShutdown() {
	p.isShutdown = true
	// Additional shutdown logic can be added here
}
