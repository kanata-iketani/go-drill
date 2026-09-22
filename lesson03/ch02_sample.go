// lesson03 ch02: 四則演算と %（int の / は切り捨て）
package main

import "fmt"

func main() {
	fmt.Println("7 + 2 =", 7+2)
	fmt.Println("7 - 2 =", 7-2)
	fmt.Println("7 * 2 =", 7*2)

	// int 同士の / は切り捨て（Python の // 相当）
	fmt.Println("7 / 2 =", 7/2)

	// 小数の答えが欲しいときは float64 に揃えてから割ります
	fmt.Println("7.0 / 2.0 =", float64(7)/float64(2))

	// % は割り算の余り（Python と同じ）
	fmt.Println("7 % 2 =", 7%2)
}
