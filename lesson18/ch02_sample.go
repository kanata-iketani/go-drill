// lesson18 ch02: 戻り値の型を書く
package main

import "fmt"

// 引数リストの後ろの int が「戻り値の型」です
func add(a, b int) int {
	return a + b
}

// 戻り値がない関数は型を書きません
func hello() {
	fmt.Println("hello")
}

func main() {
	total := add(3, 5) // 返ってきた値を変数で受け取る
	fmt.Println(total)
	fmt.Println(add(total, 10)) // 式の中でも使える
	hello()
}
