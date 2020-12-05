package skiplist


// iterate node in the Skiplist
type Iterator struct {
	list *Skiplist
	node *Node
}

// Return true iff the iterator is postiioned at a valid node
func (it *Iterator) Valid() bool {
	return it.node != nil
}

//Return the key at the current node
func (it *Iterator) Key() interface{} {
	return it.node.data
}

//Advances to the next position in the 0 level
//REQUIRES:Valid()
func (it *Iterator) Next() {
	it.list.mu.RLock()
	defer it.list.mu.RUnlock()

	it.node = it.node.getNext(0)
}

//Advances to the prev position in the 0 level
//REQUIRES:Valid()
func (it *Iterator) Prev() {
	it.list.mu.RLock()
	defer it.list.mu.RUnlock()

	it.node = it.list.FindLessThan(it.node.data)
}

// Advance to the first entry with a key >=target
func (it *Iterator) Seek(target interface{}){
	it.list.mu.RLock()
	defer it.list.mu.RUnlock()

	it.node,_=it.list.FindGreaterOrEquals(target)
}

func (it *Iterator)SeekToFirst(){
	it.list.mu.RLock()
	defer it.list.mu.RUnlock()

	it.node=it.list.head.getNext(0)
}

func (it *Iterator)SeekToLast(){
	it.list.mu.RLock()
	defer it.list.mu.RUnlock()

	it.node=it.list.FindLast()
}
