package entity

import (
	"github.com/google/uuid"
	"time"
)

// SnapshotMeta identifies network state.
type SnapshotMeta struct {
	Identifier uuid.UUID `json:"id"`
	CapturedAt time.Time `json:"capturedAt"`
	Nodes      []string  `json:"nodes"`
}

func (snapshot SnapshotMeta) ID() string {
	return snapshot.Identifier.String()
}
