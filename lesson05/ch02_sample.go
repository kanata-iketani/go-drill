// lesson05 ch02: len と s[i] はバイト単位
package main

import "fmt"

func main() {
	en := "Hello"
	ja := "こんにちは"

	// 英数字は 1 文字 1 バイトなので直感どおりです
	fmt.Println(len(en))       // 5
	fmt.Println(en[0])         // 72（'H' のバイト値）
	fmt.Println(string(en[0])) // "H" に戻すには string() が必要

	// 日本語は UTF-8 で 1 文字 3 バイトです
	fmt.Println(len(ja)) // 15（5 文字 × 3 バイト）

	// ja[0] は「こ」の 1 バイト目にすぎません
	fmt.Println(ja[0]) // 227
}
