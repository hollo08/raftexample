package cache

import (
	"encoding/json"
	"io"
	"log"

	"github.com/hashicorp/raft"
)

type FSM struct {
	ctx *RaftCachedContext
	log *log.Logger
}

type LogEntryData struct {
	Key   string
	Value string
}

// Apply applies a Raft log entry to the key-value store.
func (f *FSM) Apply(logEntry *raft.Log) interface{} {
	e := LogEntryData{}
	if err := json.Unmarshal(logEntry.Data, &e); err != nil {
		panic("Failed unmarshaling Raft log entry. This is a bug.")
	}
	ret := f.ctx.RCC.Cm.Set(e.Key, e.Value)
	f.log.Printf("fms.Apply(), logEntry:%s, ret:%v\n", logEntry.Data, ret)
	return ret
}

// Snapshot returns a latest snapshot
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return &snapshot{cm: f.ctx.RCC.Cm}, nil
}

// Restore stores the key-value store to a previous state.
func (f *FSM) Restore(serialized io.ReadCloser) error {
	return f.ctx.RCC.Cm.UnMarshal(serialized)
}
