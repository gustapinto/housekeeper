package fixtureObserver

import (
	"strings"

	"github.com/fsnotify/fsnotify"

	"housekeeper/src/fswatcher"
	"housekeeper/src/handler"
	"housekeeper/src/observer"
)

type FixtureObserver struct {
	observer.Observer
}

func (o *FixtureObserver) Observe() {
	fswatcher.MakeRecursiveFilesystemWatcher("../fixtures", o)
}

func (o *FixtureObserver) Handle(event fsnotify.Event) {
	filename := event.Name

	switch event.Op {
		case fsnotify.Create:
			switch {
			case strings.Contains(filename, ".txt"):
				handler.TxtHandler(filename)

			case strings.Contains(filename, ".py"):
				handler.PyHandler(filename)
			}
	}
}
