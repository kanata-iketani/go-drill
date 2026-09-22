// lesson18 ch04: 無名関数（lambda の代わり）
package main

import "fmt"

func main() {
	// 名前の代わりに変数に代入して使います
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(3))

	// lambda と違い、複数行のロジックも書けます
	clamp := func(x int) int {
		if x < 0 {
			return 0
		}
		if x > 100 {
			return 100
		}
		return x
	}
	fmt.Println(clamp(-5), clamp(50), clamp(300))
}
