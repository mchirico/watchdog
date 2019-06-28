package util

import (
	"github.com/mchirico/tlib/util"
	"testing"
)



func TestExampleNewWatcher(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	Wrapper()

}
