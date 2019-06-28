package util

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"testing"
	"time"
)

func TestExampleNewWatcher(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(15*time.Second))
	defer cancel()

	msg := "test here..."

	go NewWatcher(ctx, util.PWD(), func(v string) {

		fmt.Printf("thingB: %v\n", v)
		time.Sleep(1 * time.Second)
		fmt.Printf("msg: %v\n", msg)

	})

	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 10; i++ {
		file := fmt.Sprintf("junk%d", i)
		util.WriteString(file, "abc", 0644)
		time.Sleep(1 * time.Millisecond)
	}
	time.Sleep(15 * time.Second)

}
