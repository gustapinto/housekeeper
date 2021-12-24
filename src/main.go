package main

import (
	"housekeeper/src/fswatcher"
)

func main() {
	pathsToWatch := []string{
		"./",
		"../fixtures/",
	}

	fswatcher.MakeFilesystemWatchers(pathsToWatch)
}
