package util

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Wrapper() {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(5*time.Second))
	defer cancel()

	c := make(chan string)

	go ExampleNewWatcher(ctx, util.PWD(), c)
	time.Sleep(1 * time.Second)

	util.WriteString("test", "data", 0644)

	for {

		select {

		case v, ok := <-c:
			fmt.Printf("tmp: %v, %v\n", v, ok)
			if !ok {
				fmt.Printf("Channel closed\n")
				return
			}

		}
	}
}

func ExampleNewWatcher(ctx context.Context, dir string, c chan string) {
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
