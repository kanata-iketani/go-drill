// lesson01 ch03 解答
package main

import "fmt"

func main() {
	// var 変数名 型 = 値 が基本形です。型を明示することで
	// 「year は整数、version は小数」という約束がコンパイル時に固定され、
	// 間違った型の代入を実行前に検出できます。
	var year int = 2026
	var version float64 = 1.22

	fmt.Println("year:", year)
	fmt.Println("version:", version)
}
