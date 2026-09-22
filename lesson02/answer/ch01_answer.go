// lesson02 ch01 解答
package main

import "fmt"

func main() {
	fmt.Println("Go の標準出力")
	// Println にカンマで複数の値を渡すと、自動で空白区切りになります。
	// "level 5" と 1 つの文字列にしなくてよいので、
	// 文字列と数値を混ぜて出すときはこの形が最も簡単です。
	fmt.Println("level", 5)
}
