// lesson16 ch02: make と追加・更新
package main

import "fmt"

func main() {
	// 空の map は make で作ります（var だけの nil map には書き込めません）
	stock := make(map[string]int)

	// キーがなければ追加
	stock["pen"] = 10
	stock["note"] = 5

	// キーがあれば上書き（更新）
	stock["pen"] = 8

	fmt.Println(stock["pen"])
	fmt.Println(stock["note"])
	fmt.Println(len(stock))
}
