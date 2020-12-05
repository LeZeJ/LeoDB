package skiplist

import (
	"math/rand"
	"sync"
	"time"

	"github.com/LeZeJ/LeoDB/utils"
)

var maxLevel int = 32
var kBranching = 4
var random *rand.Rand

type Skiplist struct {
	level      int //有数据的最大层数
	head       *Node
	comparator utils.Comparator //比较器，用于节点data的排序
	mu         sync.RWMutex     //读写互斥锁，用于对跳表的读写互斥操作
}

//新建跳表
//参数：确定比较函数的类型
func NewSkipList(comp utils.Comparator) *Skiplist {
	// 根据当前时间生成的随机数
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
	sl := new(Skiplist)
	sl.head = newNode(nil, maxLevel)
	sl.level = 1
	sl.comparator = comp
	sl.mu = sync.RWMutex{}
	return sl
}

//获得随机高度,并刷新跳表最大高度
func (list *Skiplist) GetRandomHeight() (h int) {
	h = 1
	for ; h < maxLevel && random.Int()%kBranching == 0; h++ {
	}
	if h > list.level {
		list.level = h
	}
	return
}

//查找数据
func (list *Skiplist) Contains(data interface{}) bool {
	list.mu.RLock()
	defer list.mu.RUnlock()

	p := list.head
	for curlevel := list.level - 1; curlevel >= 0; curlevel-- {
		for {
			if p.next[curlevel] == nil || list.comparator(p.next[curlevel].data, data) > 0 {
				break //到下一层进行操作
			}
			if list.comparator(p.next[curlevel].data, data) == 0 {
				return true
			}
			p = p.next[curlevel]
		}
	}
	return false
}

//插入数据
func (list *Skiplist) Insert(data interface{}) {
	list.mu.Lock()
	defer list.mu.Unlock()

	height := list.GetRandomHeight()
	newnode := newNode(data, height)
	p := list.head
	//进行插入(这里特别关键，必须是从有数据的最上层进行查找的)
	for curlevel := list.level - 1; curlevel >= 0; curlevel-- {
		for {
			if p.next[curlevel] == nil || list.comparator(p.next[curlevel].data, data) > 0 {
				if curlevel < height {
					newnode.next[curlevel] = p.next[curlevel]
					p.next[curlevel] = newnode
				}
				break
			}
			p = p.next[curlevel]
		}
	}

}

//Delete 删除数据
func (list *Skiplist) Delete(data interface{}) (flag bool) {
	list.mu.Lock()
	defer list.mu.Unlock()

	p := list.head
	flag = false
	for curlevel := list.level-1; curlevel >= 0; curlevel-- {
		for {
			//注意判断条件的顺序
			if p.next[curlevel] == nil || list.comparator(p.next[curlevel].data, data) > 0 {
				break
			}
			//一层一层的删除节点
			if p.next[curlevel] != nil && list.comparator(p.next[curlevel].data, data) == 0 {
				//删除节点
				flag = true
				p.next[curlevel] = p.next[curlevel].next[curlevel]
				break
			}
			p = p.next[curlevel]
		}
	}
	return
}

func (list *Skiplist)NewIterator()(*Iterator){
	var it Iterator
	it.list=list
	return &it
}

func (list *Skiplist) FindGreaterOrEquals(data interface{})(next *Node,pre [32]*Node){
	
	p:=list.head
	for height:=list.level-1;height>=0;height--{
		for{
			if p.next[height]==nil||list.comparator(p.next[height].data,data)>=0{
				pre[height]=p
				if height==0{
					next=p.next[height]
					return 
				}
				break
			}
			p=p.next[height]
		}	
	}
	return 
}

func (list *Skiplist) FindLessThan(data interface{})( pre *Node){
	p:=list.head
	for height:=list.level-1;height>=0;height--{
		for{
			if p.next[height]==nil||list.comparator(p.next[height].data,data)>=0{
				if height==0{
					pre=p
					return 
				}
				break
			}
			p=p.next[height]
		}	
	}
	return 
}

func (list *Skiplist)FindLast() (last *Node){
	p:=list.head

	for curlevel:=list.level;curlevel>=0;curlevel--{
		for{
			if p.next[curlevel]==nil{
				if curlevel==0{
					last=p
					return
				}
				break
			}
			p=p.next[curlevel]
		}
		
	}
	return 
}

