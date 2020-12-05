package memtable

import (
	"github.com/LeZeJ/LeoDB/internal"
	"github.com/LeZeJ/LeoDB/skiplist"
)

//包装Iterator，使其遍历存储值为internalkey的节点
//这个细节主意好，这里是基于internalkey编码规则构建的memtable，如果变换一个
//编码规则，则skiplist原本的结构依然可以利用这么写，代码的复用性
type Iterator struct {
	listIter *skiplist.Iterator
}

func (it *Iterator) Valid() bool {
	return it.listIter.Valid()
}

//保证传进来的是internalkey的类型

func (it *Iterator) Key() *internal.InternalKey {
	//断言，若不是InternalKey类型，直接报错
	return it.listIter.Key().(*internal.InternalKey)
}

//Advances to the next position in the 0 level
//REQUIRES:Valid()
func (it *Iterator) Next() {
	it.listIter.Next()
}

//Advances to the prev position in the 0 level
//REQUIRES:Valid()
func (it *Iterator) Prev() {
	it.listIter.Prev()
}

// Advance to the first entry with a key >=target
func (it *Iterator) Seek(target interface{}) {
	it.listIter.Seek(target)
}

func (it *Iterator) SeekToFirst() {
	it.listIter.SeekToFirst()
}

func (it *Iterator) SeekToLast() {
	it.listIter.SeekToLast()
}

