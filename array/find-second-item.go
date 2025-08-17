package array_test

import (
	"errors"
	"math"
	"strconv"
)

/**
✅ 题目：查找数组中第二大的元素
实现一个函数，返回数组中第二大的数字（要求一次遍历完成）。

示例
输入：[3, 1, 4, 5, 2]
输出：4

考点
	•	一次遍历维护两个变量 max1、max2
	•	边界值判断（数组长度 < 2）
*/

/**
思考：
- 要寻找第二大的，则需要记录 第一大、第二大 两个
- 遍历时，要从头到尾遍历一遍，没办法减少遍历个数
- 边界思考：
- 若 len(a) <= 1，则没有第二大的元素
*/

func FindSecondItem(a []int) (int, error) {
	if len(a) <= 1 {
		return 0, errors.New("数组只有" + strconv.Itoa(len(a)) + "个元素，无法定位第二大的元素")
	}

	// 使用 math.MinInt 初始化 max1， max2
	max1, max2 := math.MinInt, math.MinInt

	for i := 0; i <= len(a)-1; i++ {
		if a[i] > max1 { // 找到最大值元素，则将当前最大值赋予第二大值，并更新最大值
			max2 = max1
			max1 = a[i]
		} else if a[i] < max1 && a[i] > max2 { // 找到比最大值小，还比第二大值大的元素，则更新第二大值
			max2 = a[i]
		}
	}

	if max2 == math.MinInt {
		return 0, errors.New("数组中没有第二大的元素（所有元素相等）")
	}

	return max2, nil
}
