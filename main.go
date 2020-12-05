package main

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	// sl := skiplist.NewSkipList(utils.IntComparator)
	// for i := 1; i < 100000; i++ {
	// 	sl.Insert(i)
	// }

	// fmt.Println(sl)
	// fmt.Println(sl.Contains(5515))
	// fmt.Println(sl.Delete(5515))

	// fmt.Println(sl.Contains(5515))

	// it := sl.NewIterator()
	// it.Seek(5515)
	// it.Next()
	// it.Next()
	// it.SeekToLast()
	// it.SeekToFirst()
	// fmt.Println(it.Key())

	// memTable := memtable.NewMemtable()
	// memTable.Add(1234567, internal.TypeValue, []byte("aadsa34a"), []byte("bbb3423"))
	// value, _ := memTable.Get([]byte("aadsa34a"))
	// fmt.Println(string(value))
	// fmt.Println(memTable.ApproximateMemoryUsage())
	db, err := leveldb.OpenFile("/test", nil)
	if err != nil {
		log.Fatal("open db error!")
	}
	err = db.Put([]byte("key"), []byte("value"), nil)
	if err != nil {
		log.Fatal("put db error!")
	}
	data, err := db.Get([]byte("key"), nil)
	if err != nil {
		log.Fatal("get db error!")
	}
	fmt.Println(string(data))
	defer db.Close()
}
