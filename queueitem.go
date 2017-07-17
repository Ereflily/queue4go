package queue4go

type QueueItem struct {
	data interface{}
}

func (this *QueueItem) Data() interface{} {
	return this.data
}

func NewQueueItem(data interface{}) *QueueItem {
	return &QueueItem{
		data: data,
	}
}
