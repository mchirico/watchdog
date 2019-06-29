package util

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

func NewWatcher(ctx context.Context, pwd string, fn func(v string)) {


	pipeline := newWatcher(ctx, pwd)
	for p := range pipeline {
		fn(p)
	}


}

func newWatcher(ctx context.Context, dir string) <-chan string {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}



	err = watcher.Add(dir)
	if err != nil {
		log.Fatal(err)
	}

	watchResult := make(chan string)

	go func() {
		defer watcher.Close()
		defer close(watchResult)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					watchResult <- "!ok"
					return
				}
				msg := fmt.Sprintf("event: %s", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					msg = fmt.Sprintf("%s modified file: %s\n", msg, event.Name)

				}
				watchResult <- msg
			case err, ok := <-watcher.Errors:
				if !ok {
					watchResult <- "Watcher Error: "
					return
				}
				watchResult <- fmt.Sprintf("err: %s\n",err)

			case <-ctx.Done():
				return
			}
		}
	}()



	return watchResult
}
