// lesson10 ch01: if（条件に括弧なし・{} は必須）
package main

import "fmt"

func main() {
	x := 42

	// 条件に () は付けず、{} は必ず書きます
	if x > 10 {
		fmt.Println("10より大きい")
	}

	// 処理が 1 行でも {} は省略できません
	if x%2 == 0 {
		fmt.Println("偶数")
	}

	// Python の pass に相当するのは空ブロック {} です
	if x > 100 {
	}
	fmt.Println("おわり")
}
