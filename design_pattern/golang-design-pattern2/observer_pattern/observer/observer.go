package observer

type Observer[event any] interface {
	OnNotify(handler func(msg string))
}

type Notifier[event any] interface {
	Register(o Observer[event])
	Unregister(o Observer[event])
	Notify(e event)
}
