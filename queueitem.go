package queue4go

import (
	"time"
)

type QueueItem struct {
	data     interface{}
	createAt time.Time
}

func (this *QueueItem) CreatAt() time.Time {
	return this.createAt
}

func (this *QueueItem) Data() interface{} {
	return this.data
}

func NewQueueItem(data interface{}) *QueueItem {
	t := time.Now()
	return &QueueItem{
		data:     data,
		createAt: t,
	}
}
