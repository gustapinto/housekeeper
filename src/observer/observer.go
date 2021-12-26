package observer

import (
	"github.com/fsnotify/fsnotify"
)

type Observer interface {
	Observe()
	Handle(fsnotify.Event)
}
