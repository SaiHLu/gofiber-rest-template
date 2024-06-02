package queuetype

type QueueType int

const (
	SendEmail QueueType = iota
	ResizeImage
)

func (q QueueType) String() string {
	return [...]string{"send:email", "resize:image"}[q]
}

func (q QueueType) EnumIndex() int {
	return int(q)
}
