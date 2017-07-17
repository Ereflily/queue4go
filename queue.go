package queue4go

import (
	"sync"
)

var (
	queue = make(map[string]*QueueTable)
	mutex sync.RWMutex
)

func Queue(name string) *QueueTable {
	mutex.RLock()
	_, ok := queue[name]
	mutex.RUnlock()
	if !ok {
		mutex.Lock()
		_, ok := queue[name]
		if !ok {
			queue[name] = &QueueTable{
				items:  make([]*QueueItem, 0, 8),
				length: 1000000,
			}
		}
		mutex.Unlock()
	}
	return queue[name]
}
