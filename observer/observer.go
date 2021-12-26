package observer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Observer interface {
	HandleEvent(fsnotify.Event)
	HandleError(error)
}

// Enable fsnotify to recusive watch a root directory
func watchRecursive(fswatcher *fsnotify.Watcher, path string) {
	err := filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.Mode().IsDir() {
			return fswatcher.Add(walkPath)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

// Dispatch Observer handlers on event and error
func watchFilesystemEvents(fswatcher *fsnotify.Watcher, observer Observer) {
	for {
		select {
		case event := <-fswatcher.Events:
			observer.HandleEvent(event)

		case err := <-fswatcher.Errors:
			observer.HandleError(err)
		}
	}
}

// Recursive watch a single root path and dispatch Observer handlers
func Observe(path string, observer Observer) {
	fswatcher, err := fsnotify.NewWatcher();

	if err != nil {
		log.Fatal(err)
	}

	defer fswatcher.Close()

	watchRecursive(fswatcher, path)

	done := make(chan bool)

	go watchFilesystemEvents(fswatcher, observer)

	<-done
}

// Recursive watch multiple root paths and dispatch Observer handlers
func ObserveMultiple(paths []string, observer Observer) {
	fswatcher, err := fsnotify.NewWatcher();

	if err != nil {
		log.Fatal(err)
	}

	defer fswatcher.Close()

	for _, path := range paths {
		watchRecursive(fswatcher, path)
	}

	done := make(chan bool)

	go watchFilesystemEvents(fswatcher, observer)

	<-done
}
