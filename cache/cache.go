package cache

import (
	"encoding/json"
	"io"
	"log"
	"sync"
)

type RaftCached struct {
	Opts *options
	Log  *log.Logger
	Cm   *Manager
	Raft *RaftNodeInfo
}

type RaftCachedContext struct {
	RC *RaftCached
}

type Manager struct {
	data map[string]string
	sync.RWMutex
}

func NewCacheManager() *Manager {
	cm := &Manager{}
	cm.data = make(map[string]string)
	return cm
}

func (c *Manager) Get(key string) string {
	c.RLock()
	ret := c.data[key]
	c.RUnlock()
	return ret
}

func (c *Manager) Set(key string, value string) error {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
	return nil
}

// Marshal serializes cache data
func (c *Manager) Marshal() ([]byte, error) {
	c.RLock()
	defer c.RUnlock()
	dataBytes, err := json.Marshal(c.data)
	return dataBytes, err
}

// UnMarshal deserializes cache data
func (c *Manager) UnMarshal(serialized io.ReadCloser) error {
	var newData map[string]string
	if err := json.NewDecoder(serialized).Decode(&newData); err != nil {
		return err
	}
	c.Lock()
	defer c.Unlock()
	c.data = newData
	return nil
}
