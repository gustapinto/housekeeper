package fswatcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"housekeeper/src/handler"
)

func Watch(done chan bool, watcher *fsnotify.Watcher) {
	for {
		// A select statement works like a switch but handling
		// channel operations instead of normal values
		select {
		// Get fsnotify events and dispatch workers to handle
		case event, _ := <-watcher.Events:
			operation := event.Op

			// Check if it is a creation event and
			if operation == fsnotify.Create {
				filename := event.Name

				switch {
				case strings.Contains(filename, ".txt"):
					handler.TxtHandler(filename)

				case strings.Contains(filename, ".py"):
					handler.PyHandler(filename)
				}
			}
		// Check for fsnotify error events
		case err, _ := <-watcher.Errors:
			log.Fatal(err)
		}
	}
}

func MakeFilesystemWatcher(path string) {
	watcher, err := fsnotify.NewWatcher()

	defer watcher.Close()

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go Watch(done, watcher)

	watcher.Add(path)

	<-done
}

func MakeRecursiveFilesystemWatcher(path string) {
	watcher, err := fsnotify.NewWatcher()

	defer watcher.Close()

	if err != nil {
		log.Fatal(err)
	}

	// Add recursion like feature to watcher by walking over the
	// directories inside the desired path and adding they to the watcher
	// list
	filepath.Walk(path, func(p string, fi os.FileInfo, err error) error {
		if fi.Mode().IsDir() {
			return watcher.Add(p)
		}

		return nil
	})

	done := make(chan bool)

	go Watch(done, watcher)

	<-done
}
