package array_test

import (
	"fmt"
)

/**
✅ 题目：数组值传递与引用语义对比

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
*/

/**
答：

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

func modifyArray(arr [3]int) {
	arr[0] = 100
}

func modifySlice(s []int) {
	s[0] = 100
}

// 值传递和引用传递对比
func ValuePassingVsReferenceSemantics() {
	a := [3]int{1, 2, 3}
	s := []int{1, 2, 3}

	modifyArray(a)
	modifySlice(s)

	fmt.Println("a:", a)
	fmt.Println("s:", s)
}
