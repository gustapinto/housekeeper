package fswatcher

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"housekeeper/src/observer"
)

func Watch(done chan bool, watcher *fsnotify.Watcher, o observer.Observer) {
	for {
		// A select statement works like a switch but handling
		// channel operations instead of normal values
		select {
		// Get fsnotify events and dispatch workers to handle
		case event, _ := <-watcher.Events:
			o.Handle(event)

		// Check for fsnotify error events
		case err, _ := <-watcher.Errors:
			log.Fatal(err)
		}
	}
}

func MakeFilesystemWatcher(path string, o observer.Observer) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)

	go Watch(done, watcher, o)

	watcher.Add(path)

	<-done
}

func MakeRecursiveFilesystemWatcher(path string, o observer.Observer) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	absolutePath, err := filepath.Abs(path)

	if err != nil {
		log.Fatal(err)
	}

	// Add recursion like feature to watcher by walking over the
	// directories inside the desired path and adding they to the watcher
	// list
	filepath.Walk(absolutePath, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if fi.Mode().IsDir() {
			return watcher.Add(p)
		}

		return nil
	})

	done := make(chan bool)

	go Watch(done, watcher, o)

	<-done
}
