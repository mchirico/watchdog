package util

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

func NewWatcher(ctx context.Context, pwd string, fn func(v string)) {

	c := make(chan string)

	go newWatcher(ctx, pwd, c)

	for {

		select {

		case v, ok := <-c:
			fn(v)

			if !ok {
				fmt.Printf("Channel closed\n")
				return
			}

		}
	}
}

func newWatcher(ctx context.Context, dir string, c chan string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	cWatcherExit := make(chan string)

	go func() {

		defer close(c)
		defer close(cWatcherExit)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					c <- "!ok"
					return
				}
				msg := fmt.Sprintf("event: %s", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					msg = fmt.Sprintf("%s modified file: %s\n", msg, event.Name)

				}
				c <- msg
			case err, ok := <-watcher.Errors:
				if !ok {
					c <- "Watcher Error: "
					return
				}
				c <- "error"
				log.Println("error:", err)

			case <-ctx.Done():
				c <- "ctx Done "
				return
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	<-cWatcherExit
}
