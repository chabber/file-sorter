package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	ext := filepath.Ext(fn)
	fmt.Println("ext: ", ext)

	switch ext {
	case ".jpeg":
		fmt.Println("moving file")
		err := os.Rename(fn, "temp.jpeg")

		if err != nil {
			fmt.Println("error moving file: ", err)
		}
	}
}
