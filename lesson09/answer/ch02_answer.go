// lesson09 ch02 解答
package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Python の x, y = y, x と同じ構文です。右辺が先にすべて評価されるため、
	// 一時変数を用意しなくても正しく入れ替わります。
	x, y = y, x

	fmt.Println(x)
	fmt.Println(y)
}
