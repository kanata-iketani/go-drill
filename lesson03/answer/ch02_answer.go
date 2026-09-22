// lesson03 ch02 解答
package main

import "fmt"

func main() {
	// int 同士の / は自動的に切り捨てになるため、
	// Python の 100 // 7 のような専用演算子は不要です。
	// 「何個ずつ・何個余る」の組は / と % のセットで求めるのが定石です。
	fmt.Println("1人あたり", 100/7, "個")
	fmt.Println("余り", 100%7, "個")
}
