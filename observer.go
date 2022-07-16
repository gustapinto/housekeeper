package observer

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Observer is the basic interface used on housekeeper functions
type Observer interface {
	// HandleEvent handles fsnotify events
	HandleEvent(fsnotify.Event)
	// HandleError handles fsnotify event errors
	HandleError(error)
}

// Observe Recursive watch a single root path and dispatch Observer handlers
func Observe(path string, observer Observer) {
	fswatcher := newFsnotifyWatcher()
	defer fswatcher.Close()

	addPathRecursive(fswatcher, path)
	watchFilesystemEvents(fswatcher, observer)
}

// ObserveMultiple Recursive watch multiple root paths and dispatch Observer handlers
func ObserveMultiple(paths []string, observer Observer) {
	fswatcher := newFsnotifyWatcher()
	defer fswatcher.Close()

	for _, path := range paths {
		addPathRecursive(fswatcher, path)
	}

	watchFilesystemEvents(fswatcher, observer)
}

func newFsnotifyWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	return watcher
}

// addPathRecursive Enables recusive watch on a root directory
func addPathRecursive(fswatcher *fsnotify.Watcher, root string) {
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return nil
		}

		fswatcher.Add(path)

		return nil
	}); err != nil {
		panic(err)
	}
}

// watchFilesystemEvents Dispatch Observer handlers on event and error
func watchFilesystemEvents(fswatcher *fsnotify.Watcher, observer Observer) {
	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-fswatcher.Events:
				observer.HandleEvent(event)
			case err := <-fswatcher.Errors:
				observer.HandleError(err)
			}
		}
	}()

	<-done
}
