package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go monitorDir(watcher)

	err = watcher.Add("/Users/jhunt/Downloads")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func monitorDir(watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			switch event.Op {
			case fsnotify.Create:
				moveFile(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}

func moveFile(fn string) {

}
