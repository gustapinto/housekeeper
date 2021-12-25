package main

import (
	"flag"

	"housekeeper/src/fswatcher"
)

func main() {
	path := flag.String("p", "../fixtures", "Define a path to watch")
	recursive := flag.Bool("r", false, "Define recursion on the watched path")

	flag.Parse()

	if *recursive {
		fswatcher.MakeRecursiveFilesystemWatcher(*path)
	} else {
		fswatcher.MakeFilesystemWatcher(*path)
	}
}
