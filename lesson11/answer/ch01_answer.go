// lesson11 ch01 解答
package main

import "fmt"

func main() {
	age := 18

	// 比較演算子の結果はそのまま bool 型の値になるので、
	// 変数に入れてから出力しても、式を直接 Println に渡しても同じです。
	ok := age >= 20
	fmt.Println(ok)
	fmt.Println(age == 18)
}
