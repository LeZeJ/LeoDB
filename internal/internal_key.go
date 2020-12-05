package internal

//internalKey内部变量和byte数组的转换
import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

type ValueType int8

const (
	TypeDeletion ValueType = 0
	TypeValue    ValueType = 1
)

type InternalKey struct {
	Seq       uint64 //时间戳
	Type      ValueType //操作类型，为0的话代表删除
	//用户的key值和value值用字节的方式存储
	UserKey   []byte
	UserValue []byte
}

//NewInternalKey 构建一个InternalKey
func NewInternalKey(seq uint64, valueType ValueType, key, value []byte) *InternalKey {
	var internalKey InternalKey
	internalKey.Seq = seq
	internalKey.Type = valueType
	internalKey.UserKey = make([]byte, len(key))
	copy(internalKey.UserKey, key)
	internalKey.UserValue = make([]byte, len(value))
	copy(internalKey.UserValue, value)

	return &internalKey

}

//EncodeTo 将Internalkey中的值转换为byte数组
func (key *InternalKey) EncodeTo(w io.Writer) error {
	//将internalkey中的值以小端的方式写进w的字节流中
	binary.Write(w, binary.LittleEndian, key.Seq)
	binary.Write(w, binary.LittleEndian, key.Type)
	binary.Write(w, binary.LittleEndian, int32(len(key.UserKey)))
	binary.Write(w, binary.LittleEndian, key.UserKey)
	binary.Write(w, binary.LittleEndian, int32(len(key.UserValue)))
	return binary.Write(w, binary.LittleEndian, key.UserValue)
}

func (key *InternalKey) DecodeFrom(r io.Reader) error {
	var tmp int32
	//从r的字节流中读出Internalkey的值，读到响应的地址中
	binary.Read(r, binary.LittleEndian, &key.Seq)
	binary.Read(r, binary.LittleEndian, &key.Type)
	binary.Read(r, binary.LittleEndian, &tmp)
	key.UserKey = make([]byte, tmp)
	binary.Read(r, binary.LittleEndian, key.UserKey)
	binary.Read(r, binary.LittleEndian, &tmp)
	key.UserValue = make([]byte, tmp)
	return binary.Read(r, binary.LittleEndian, key.UserValue)
}

//InternalKeyComparator internalkey的比较函数，相同的key值，序列号大的小
func InternalKeyComparator(a, b interface{}) int {
	aKey := a.(*InternalKey)
	bKey := b.(*InternalKey)
	r := UserKeyComparator(aKey.UserKey, bKey.UserKey)
	//如果key值相同，则比较序列号
	if r == 0 {
		aNum := aKey.Seq
		bNum := bKey.Seq
		if aNum > bNum {
			r = -1
		} else {
			r = +1
		}
	}
	return r
}
// 字节数组的大小比较
func UserKeyComparator(a, b interface{}) int {
	aKey := a.([]byte)
	bKey := b.([]byte)
	//字节数组的比较函数
	return bytes.Compare(aKey, bKey)
}

//通过给定的key值构建一个internalKey
func LookupKey(key []byte) *InternalKey {
	return NewInternalKey(math.MaxUint64, TypeValue, key, nil)
}
