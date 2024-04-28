package entity

import (
	"github.com/google/uuid"
	"time"
)

// SnapshotMeta identifies network state.
type SnapshotMeta struct {
	Identifier uuid.UUID `json:"id"`
	CapturedAt time.Time `json:"capturedAt"`
	Nodes      []string
	Links      []string
}

func (snapshot SnapshotMeta) ID() string {
	return snapshot.Identifier.String()
}
