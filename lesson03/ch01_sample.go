// lesson03 ch01: int と float64 は混ぜられない
package main

import "fmt"

func main() {
	n := 3   // int
	f := 1.5 // float64

	fmt.Println("n =", n)
	fmt.Println("f =", f)

	// fmt.Println(n + f) // コンパイルエラー: int + float64 はできません

	// float64 に揃えてから計算します
	fmt.Println("n + f =", float64(n)+f)

	// 逆に int に揃えると小数部分は切り捨てられます
	fmt.Println("n + int(f) =", n+int(f))
}
