package watcher

type Watcher interface {
	Start() error
	Reload()
	Stop()
}
