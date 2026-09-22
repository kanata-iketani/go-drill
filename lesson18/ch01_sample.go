// lesson18 ch01: func と引数の型
package main

import "fmt"

// greet は名前と回数を受け取ってあいさつを出力します。
// 引数には名前のあとに必ず型を書きます（name string, times int）
func greet(name string, times int) {
	for range times {
		fmt.Println("Hello,", name)
	}
}

// 同じ型が続くときは a, b int とまとめて書けます
func printSum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	greet("Gopher", 2)
	printSum(3, 5)
}
