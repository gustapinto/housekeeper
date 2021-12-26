package observer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Observer interface {
	Observe()
	HandleEvent(fsnotify.Event)
	HandleError(error)
}

func getRecursiveDirs(fswatcher *fsnotify.Watcher, path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Fatal(err)
	}

	if info.Mode().IsDir() {
		return fswatcher.Add(path)
	}

	return nil
}

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

func Watch(path string, observer Observer) {
	fswatcher, err := fsnotify.NewWatcher();

	if err != nil {
		log.Fatal(err)
	}

	defer fswatcher.Close()

	err = filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
		return getRecursiveDirs(fswatcher, walkPath, info, err)
	})

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go watchFilesystemEvents(fswatcher, observer)

	<-done
}
