# Housekeeper

A tiny wrapper to create [fsnotify](https://github.com/fsnotify/fsnotify) recursive watchers with a framework like look and higher abstraction level sintax

## Example usage
```go
package main

import (
	"log"

	"github.com/gustapinto/housekeeper/observer"
	"github.com/fsnotify/fsnotify"
)

type DotObserver struct{}

func (o *DotObserver) HandleEvent(event fsnotify.Event) {
	switch event.Op {
		case fsnotify.Create:
			log.Print("Created: ", event.Name)
	}
}

func (o *DotObserver) HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	observer.Observe(".", &DotObserver{})
}

```
