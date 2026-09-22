// lesson01 ch03: 静的型付け
// 変数の型はコンパイル時に固定され、あとから変えられません
package main

import "fmt"

func main() {
	var age int = 20
	var height float64 = 171.5

	fmt.Println("age:", age)
	fmt.Println("height:", height)

	// age = "twenty" // これを有効にするとコンパイルエラーになります

	age = 21 // 同じ int なら再代入できます
	fmt.Println("来年の age:", age)
}
