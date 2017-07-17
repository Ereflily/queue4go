package queue4go

import (
	"sync"
)

type QueueTable struct {
	sync.RWMutex
	items  []*QueueItem
	length int
}

func (this *QueueTable) Length() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.items)
}

func (this *QueueTable) Pop() interface{} {
	this.Lock()
	defer this.Unlock()
	if len(this.items) <= 0 {
		return false
	}
	value := this.items[0].Data()
	this.items = this.items[1:]
	return value
}

func (this *QueueTable) SetMaxLength(n int) {
	this.Lock()
	defer this.Unlock()
	this.length = n
}

func (this *QueueTable) GetMaxLength() int {
	this.RLock()
	defer this.RUnlock()
	return this.length
}

func (this *QueueTable) Push(item interface{}) bool {
	data := NewQueueItem(item)
	this.Lock()
	defer this.Unlock()
	if len(this.items) > 0 && len(this.items) < this.length {
		this.items = append(this.items, data)
		return data == this.items[0]
	} else if len(this.items) == 0 {
		this.items = append(this.items, data)
	} else {
		return false
	}
	return data == this.items[0]
}

func (this *QueueTable) Pos(n int) interface{} {
	this.RLock()
	defer this.RUnlock()
	if n > len(this.items) || n < 0 {
		return false
	}

	return this.items[n].Data()
}
