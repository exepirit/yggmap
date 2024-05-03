package entity

import (
	"encoding/base64"
	"github.com/oklog/ulid/v2"
	"time"
)

// SnapshotMeta identifies network state.
type SnapshotMeta struct {
	Identifier ulid.ULID
	CapturedAt time.Time `json:"capturedAt"`
	Nodes      []string  `json:"nodes"`
}

func (snapshot SnapshotMeta) ID() string {
	return base64.StdEncoding.EncodeToString(snapshot.Identifier.Bytes())
}
