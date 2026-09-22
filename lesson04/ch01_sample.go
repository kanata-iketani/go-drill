// lesson04 ch01: var による宣言
package main

import "fmt"

func main() {
	// 基本形: var 変数名 型 = 値
	var name string = "gopher"
	var age int = 13

	// 初期値から型が明らかなら型を省略できます
	var lang = "Go"

	// 値を入れないとゼロ値（int は 0）が入ります
	var count int

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("lang:", lang)
	fmt.Println("count:", count)
}
