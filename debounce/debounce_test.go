package debounce

import (
	"fmt"
	"testing"
)

func TestTicketStore(t *testing.T) {

	tk := NewTicketStore(5)
	tk.Put("test")
	fmt.Printf("here: %v\n", tk.GetDone())
	fmt.Printf("done: %d\n", *tk.done)

	tk.Put("test1")

	fmt.Printf("here: %v\n", tk.GetDone())

}
