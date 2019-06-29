package util

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"

	"strings"
	"testing"
	"time"
)

func Write(n int) {
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < n; i++ {
		file := fmt.Sprintf("junk%d", i)
		util.WriteString(file, "abc", 0644)
		time.Sleep(100 * time.Millisecond)
	}

}

func TestExampleNewWatcher(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(1*time.Second))
	defer cancel()

	num := 10
	go Write(num)

	count0 := 0
	count1 := 1

	pipeline := Watcher(ctx, util.PWD())
	for p := range pipeline {

		if strings.Contains(p, "CREATE") {
			count0 += 1
		}
		if strings.Contains(p, "CHMOD") {
			count1 += 1
		}

	}

	if count0 != count1 && count0 < num {
		t.FailNow()
	}

}
