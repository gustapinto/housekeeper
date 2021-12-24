package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	fs "github.com/fsnotify/fsnotify"
)

func PyHandler(script string) {
	out, _ := exec.Command("python", script).Output()

	fmt.Println(out)
}

func TxtHandler(filePath string) {
	// Read a file content using the os package, this is not so
	// efficient on larger files but is a pratical and high level way
	// of doing it
	content, _ := os.ReadFile(filePath)

	fmt.Println(string(content))
}

func HandleFilesystemWatcherEvents(done chan bool, watcher *fs.Watcher) {
	for {
		// A select statement works like a switch but handling
		// channel operations instead of normal values
		select {
		// Get fsnotify events and dispatch workers to handle
		case event, _ := <-watcher.Events:
			operation := event.Op

			// Check if it is a creation event and
			if operation == fs.Create {
				filename := event.Name

				switch {
				case strings.Contains(filename, ".txt"):
					TxtHandler(filename)

				case strings.Contains(filename, ".py"):
					PyHandler(filename)
				}
			}
		// Check for fsnotify error events
		case err, _ := <-watcher.Errors:
			fmt.Println(err)
		}
	}
}

func MakeFilesystemWatchers(paths []string) {
	watcher, _ := fs.NewWatcher()

	defer watcher.Close()

	done := make(chan bool)

	// Opens a goroutine to handle file watching
	go HandleFilesystemWatcherEvents(done, watcher)

	for _, path := range paths {
		watcher.Add(path)
	}

	<-done
}

func main() {
	pathsToWatch := []string{
		"./",
		"../fixtures/",
	}

	MakeFilesystemWatchers(pathsToWatch)
}
