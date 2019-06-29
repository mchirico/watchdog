package util

import (
	"context"
	"fmt"
	"github.com/mchirico/tlib/util"
	"io/ioutil"
	"sync"

	"strings"
	"testing"
	"time"
)

func Write(n int) {
	time.Sleep(2 * time.Second)

	for i := 0; i < n; i++ {
		file := fmt.Sprintf("junk%d", i)
		util.WriteString(file, "abc", 0644)
		time.Sleep(3 * time.Millisecond)
		util.AppendString(file, " more data")
		time.Sleep(3 * time.Millisecond)
	}

}

func WriteAppend(n int) {

	for i := 0; i < n; i++ {
		file := fmt.Sprintf("junk%d", i)
		util.WriteString(file, "abc", 0644)

	}

	time.Sleep(1 * time.Second)

	for i := 0; i < n; i++ {
		file := fmt.Sprintf("junk%d", i)
		time.Sleep(3 * time.Millisecond)
		util.AppendString(file, " more data")

	}

}

func TestExampleNewWatcher(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(4*time.Second))
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

		fmt.Printf("%v\n", p)
	}

	if count0 != count1 && count0 < num {
		t.FailNow()
	}

}

type F struct {
	sync.Mutex
	file  string
	count uint64
}

func (f *F) Fn(s string) {
	f.Lock()
	defer f.Unlock()
	f.count += 1

	sarray := strings.Split(s, "\"")
	fmt.Printf("%v\n", sarray[1])
}

func (f *F) GetCount() uint64 {
	f.Lock()
	defer f.Unlock()
	return f.count
}

func (f *F) Read(s string) {
	f.Lock()
	defer f.Unlock()

	f.count += 1
	sarray := strings.Split(s, "\"")

	dat, err := ioutil.ReadFile(sarray[1])
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat), "\n")

}

func TestMonitor(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(4*time.Second))
	defer cancel()

	f := F{}

	num := 100
	go Write(num)

	Monitor(ctx, util.PWD(), "CHMOD", f.Fn)

	fmt.Printf("f.GetCount: %d\n", f.GetCount())

}

func TestMonitorRead(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(4*time.Second))
	defer cancel()

	f := F{}

	num := 10
	go WriteAppend(num)

	Monitor(ctx, util.PWD(), "CHMOD", f.Read)

	fmt.Printf("f.GetCount: %d\n", f.GetCount())

}
