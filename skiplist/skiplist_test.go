package skiplist

import (
	"fmt"
	"testing"

	"github.com/LeZeJ/LeoDB/utils"
)

func Test_Insert(t *testing.T) {
	skiplist := NewSkipList(utils.IntComparator)
	for i := 0; i < 10000; i++ {
		skiplist.Insert(i)
	}
	it:=skiplist.NewIterator()
	it.Seek(1000)
	fmt.Println(it.Key())
	it.Prev()
	fmt.Println(it.node.data)
	// fmt.Println("插入成功！")
}
