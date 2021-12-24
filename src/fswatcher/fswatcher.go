package fswatcher

import (
	"fmt"
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
			fmt.Println(err)
		}
	}
}

func MakeFilesystemWatchers(paths []string) {
	watcher, _ := fsnotify.NewWatcher()

	defer watcher.Close()

	done := make(chan bool)

	// Opens a goroutine to handle file watching
	go Watch(done, watcher)

	for _, path := range paths {
		watcher.Add(path)
	}

	<-done
}
