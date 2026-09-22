// lesson10 ch03: 条件の前の1文（if v := f(); v > 0）
package main

import "fmt"

func double(n int) int {
	return n * 2
}

func main() {
	// 条件の前に 1 文書けます。v はこの if / else の中だけで有効です
	if v := double(21); v > 40 {
		fmt.Println("40より大きい:", v)
	}

	name := "gopher"
	if l := len(name); l >= 6 {
		fmt.Println("6文字以上です")
	} else {
		fmt.Println("短い名前です:", l)
	}
	// ここでは v も l も使えません（スコープの外）
}
