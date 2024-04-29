package entity

import (
	"time"
)

// SnapshotMeta identifies network state.
type SnapshotMeta struct {
	CapturedAt time.Time `json:"capturedAt"`
	Nodes      []string  `json:"nodes"`
}

func (snapshot SnapshotMeta) ID() string {
	return snapshot.CapturedAt.UTC().Format(time.RFC3339)
}
