package memtable

import (
	"github.com/LeZeJ/LeoDB/internal"
	"github.com/LeZeJ/LeoDB/skiplist"
)

type Memtable struct {
	table       *skiplist.Skiplist
	memtableCap uint64
}

//用internalkey的比较规则构建的内存跳表结构
func New() *Memtable {
	var memTable Memtable
	memTable.table = skiplist.NewSkipList(internal.InternalKeyComparator)
	return &memTable
}

func (memTable *Memtable) NewIterator() *Iterator {
	return &Iterator{listIter: memTable.table.NewIterator()}
}

func (memTable *Memtable) Add(seq uint64, valueType internal.ValueType, key, value []byte) {
	internalKey := internal.NewInternalKey(seq, valueType, key, value)
	//字节数:序列号+type总共8个字节，key和value的长度8个字节
	memTable.memtableCap += uint64(16 + len(key) + len(value))
	memTable.table.Insert(internalKey)
}

func (memTable *Memtable) Get(key []byte) ([]byte, error) {
	lookupKey := internal.LookupKey(key)
	//使用迭代器进行查找
	it := memTable.table.NewIterator()
	//构建internalKey值并将seq值设置最大用于检索最新的具有相同Key值的internalKey
	it.Seek(lookupKey)
	if it.Valid() {
		internalKey := it.Key().(*internal.InternalKey)
		if internal.UserKeyComparator(key, internalKey.UserKey) == 0 {
			//判断valueType类型
			if internalKey.Type == internal.TypeValue {
				return internalKey.UserValue, nil
			}
			return nil, internal.ErrDeletion

		}
	}
	//没有发现相关的internalKey
	return nil, internal.ErrorNotFound
}

func (memTable *Memtable) ApproximateMemoryUsage() uint64 {
	return memTable.memtableCap
}
