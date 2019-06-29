package util

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"testing"
	"time"
)


func Write() {
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 10; i++ {
		file := fmt.Sprintf("junk%d", i)
		util.WriteString(file, "abc", 0644)
		time.Sleep(100 * time.Millisecond)
	}

}

func TestExampleNewWatcher(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(15*time.Second))
	defer cancel()


	go Write()

	pipeline := newWatcher(ctx,util.PWD())
	for p := range pipeline {
		fmt.Printf("%v\n",p)
	}




}
