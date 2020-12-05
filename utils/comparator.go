package utils

//    <0 , if a < b
//    =0 , if a == b
//    >0 , if a > b
type Comparator func(a,b interface{}) int
//IntComparator 依据int类型的比较器
func IntComparator(a,b interface{})int{
	aInt:=a.(int) //断言
	bInt:=b.(int)
	return aInt-bInt
}
