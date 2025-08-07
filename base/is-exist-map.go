package base

import "fmt"

/**
✅ 题目二：map 键存在性判断

m := map[string]int{
	"apple":  5,
	"banana": 0,
}

v1 := m["apple"]
v2 := m["banana"]
v3 := m["cherry"]

fmt.Println("v1:", v1)
fmt.Println("v2:", v2)
fmt.Println("v3:", v3)

// 如何判断 "banana" 键是存在且值为 0？

❓输出什么？怎样判断键是否存在？
fmt.Println("v1:", v1)	输出： v1: 5
fmt.Println("v2:", v2)	输出： v2: 0
fmt.Println("v3:", v3)	输出： v3: 0

对于 map 数据类型，是通过 value, ok := map[key] 的方式，通过 第二个参数来判断键是否存在的。
而键不存在时，使用 map[key] 获取值，会返回“类型零值”

*/

func IsExistOfMap() {
	m := map[string]int{
		"apple":  5,
		"banana": 0,
	}

	v1 := m["apple"]
	v2 := m["banana"]
	v3 := m["cherry"]

	fmt.Println("v1:", v1) // 输出： v1: 5
	fmt.Println("v2:", v2) // 输出： v2: 0
	fmt.Println("v3:", v3) // 输出： v3: 0

	_, ok_apple := m["apple"]
	fmt.Println("m[\"apple\"] is existed? ", ok_apple) // 输出：m["apple"] is existed?  true

	_, ok_cherry := m["cherry"]
	fmt.Println("m[\"cherry\"] is existed? ", ok_cherry) // 输出：m["cherry"] is existed?  false
}
