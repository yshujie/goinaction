package slice_test

/**
题目：切片数据去重
编写一个就地处理（in-place）的函数，用于去除 []string 切片中相邻的重复字符串元素。

要求：
1. 函数应直接在原切片上进行修改，不分配新的切片存储结果（允许底层数组复用）。
2. 只去除**相邻位置**上值相同的元素，不要求去除所有重复项。
3. 函数返回去重后的有效切片（长度可能变短）。
4. 时间复杂度要求为 O(n)，空间复杂度 O(1)。

示例：
输入：["a", "a", "b", "b", "c", "a", "a"]
输出：["a", "b", "c", "a"]

考点：
- 顺序遍历切片
- 就地修改切片并返回新长度
- 理解相邻重复与全局去重的区别
*/

/*
*
思考：
- 就地操作，即在原切片上，使用的方法是在一次循环中，维护两个队列：
- 队列一：读队列，从左向右依次读取切片元素，并将“读取的元素”和“待比较元素”进行对比
- 队列二：写队列，慢于读队列进度，只有发现需要替换的元素时才会更新

  - 读取元素：s[readerIndex]
  - 待比较元素：s[writerIndex]
  - 比较：
    读取元素 == 待比较元素，则继续读取，不更新数据
    读取元素 != 待比较元素，则更新写队列数据、更新 writerIndex
*/
func Deduplication[T comparable](s []T) []T {
	if len(s) <= 1 {
		return s
	}

	// 写队列，从第一个元素开始
	writerIndex := 0

	// 对队列，从第二个开始读取
	for readerIndex := 1; readerIndex <= len(s)-1; readerIndex++ {
		if s[readerIndex] == s[writerIndex] {
			continue
		}

		// 发现与当前元素不同的元素，则向前走一步，写入队列
		writerIndex++
		s[writerIndex] = s[readerIndex]
	}

	// 由于切片的截取规则是左闭右开，所以需要截取到 writerIndex+1 位置
	return s[:writerIndex+1]
}
