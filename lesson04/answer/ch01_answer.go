// lesson04 ch01 解答
package main

import "fmt"

func main() {
	// var 変数名 型 = 値 の基本形で宣言します。
	// 型を書いておくと「population は必ず整数」という約束が
	// コンパイル時に検査されるため、意図しない値の混入を防げます。
	var city string = "Tokyo"
	var population int = 14000000

	fmt.Println("city:", city)
	fmt.Println("population:", population)
}
