package memtable

import (
	"fmt"
	"testing"

	"github.com/LeZeJ/LeoDB/internal"
)

func Test_memtable(t *testing.T) {
	memTable := New()
	memTable.Add(1234567, internal.TypeValue, []byte("dgadga"), []byte("dhaskhd"))
	value, _ := memTable.Get([]byte("dgadga"))
	fmt.Println(value)
	fmt.Println(memTable.ApproximateMemoryUsage())
}
