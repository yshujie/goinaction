package base

import (
	"fmt"
	"reflect"
)

/**
✅ 题目三：数组值传递与引用语义对比

func modifyArray(arr [3]int) {
	arr[0] = 100
}

func modifySlice(s []int) {
	s[0] = 100
}

a := [3]int{1, 2, 3}
s := []int{1, 2, 3}

modifyArray(a)
modifySlice(s)

fmt.Println("a:", a)
fmt.Println("s:", s)

❓问：
为什么 a[0] 没有变，而 s[0] 变了？

在 go 语言中，所有的函数参数都是值传递。
也就是说，在调用 modifyArray(a) 时，是将数组 a copy 了一份，传入 modifyArray() 函数内部，
而 modifyArray() 函数内部对函数的操作，并不会影响到函数外的数组；

在调用 modifySlice(s) 时，虽然也是值传递，但 s 本身就是切片，是一个表头、描述符，
传入 modifySlice() 函数内部的切片，虽然也是值传递，但两个切片所对应的底层数组是相同的，
所以在 modifySlice() 函数内部更新切片的值，会影响函数外切片的值。

再思考一下，如果在 modifySlice() 函数内部使用 append()，对切片进行追加元素操作，
则有可能造成切边扩容，扩容后的底层数组就不再和函数外的底层数组相同了，这时候操作函数内部的切片，不会影响函数外部的切片

如下：
func modifySlice(s []int) {
	append(s, 1,2,3,4,5,6)
	s[0] = 100
}
此时 s[0] = 100 只更改函数内部切片的值，而不会影响函数外部的切片
*/

// NewArray 创建数组
func NewArray() {
	defer fmt.Println()

	// 仅声明 不初始化，此时 go 运行时会给 a0 开辟内存空间，并给每个元素赋“类型零值”
	var a01 [6]int
	var a02 [5]string
	var a03 [5]bool
	fmt.Printf("a01 type: %T; a01 value: %v \n", a01, a01)
	fmt.Printf("a02 type: %T; a02 value: %v \n", a02, a02)
	fmt.Printf("a03 type: %T; a03 value: %v \n", a03, a03)

	// 声明+初始化方式创建数组
	// a10 := [5]int{}
	a10 := [5]int{1, 2, 3, 4, 5}
	a10[0] = 10

	fmt.Printf("a10 type: %T; a10 value: %v \n", a10, a10)

	a2 := new([5]int)
	fmt.Printf("a2 type: %T; a2 value: %v \n", a2, a2)

	reflect.TypeOf(a01)
	fmt.Printf("reflect.TypeOf(a01) == reflect.TypeOf(a10): %v \n", reflect.TypeOf(a01) == reflect.TypeOf(a10))
	// fmt.Printf("a01 == a10: %v \n", a01 == a10)
}
