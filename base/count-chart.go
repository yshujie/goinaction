package base

import (
	"fmt"
	"unicode"
)

/**
✅ 题目六：使用 map 实现字符串中每个字符出现的次数统计
输入: "gopher"
输出: map[char]int -> g:1, o:1, p:1, h:1, e:1, r:1

✍ 要求
- 使用 map[rune]int 实现
- 不区分大小写，可扩展支持中文
*/

/**
思路：map 本身就是 key 唯一的集合
遍历字符串，然后每遍历到一个元素，则在集合中加入该元素，并计数+1
另外，因为获取 map 元素时，若 map 中没有，则 value 为类型零值，这就更方便了，在第一次出现时 和 第 n 次出现时处理一样，都是 +1

此外：
- 不区分大小写，则在处理元素前，统一使用小写
- 可支持中文扩展，则使用 rune ，而不是 byte
*/

func CountChart(s string) {
	chartSet := []rune{}
	chartMap := make(map[rune]int)

	for _, v := range s {
		// 统一使用小写
		v = unicode.ToLower(v)

		// 为字符计数
		chartMap[v] = chartMap[v] + 1

		// 若字符是第一次出现，则将字符追加进 set 集合
		if chartMap[v] == 1 {
			chartSet = append(chartSet, v)
		}
	}

	fmt.Print(" map[char]int ->")
	// 打印每个字符出现的次数
	for _, v := range chartSet {
		fmt.Printf("%c:%d, ", v, chartMap[v])
	}
	fmt.Println()
}
