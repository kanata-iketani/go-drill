// lesson16 ch03: カンマ ok イディオム
package main

import "fmt"

func main() {
	prices := map[string]int{"coffee": 480, "tea": 400}

	// 存在しないキーはエラーにならず、ゼロ値（int なら 0）が返ります
	fmt.Println(prices["juice"])

	// 「0 円」なのか「メニューにない」のかは、カンマ ok で区別します
	if v, ok := prices["coffee"]; ok {
		fmt.Println("coffee は", v, "円")
	}

	// 値が不要で存在だけ知りたいときは v を _ で捨てます
	if _, ok := prices["juice"]; !ok {
		fmt.Println("juice はメニューにありません")
	}
}
