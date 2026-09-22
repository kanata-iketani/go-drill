// lesson16 ch01: map リテラルと取得
package main

import "fmt"

func main() {
	// map[キーの型]値の型{...} がリテラルです
	// 複数行で書くときは最後の要素にもカンマが必要です
	prices := map[string]int{
		"apple":  120,
		"banana": 90,
		"orange": 60,
	}

	// 取得は Python の辞書と同じ書き方です
	fmt.Println(prices["apple"])
	fmt.Println(prices["orange"])

	// 取得した値は式の中でもそのまま使えます
	fmt.Println(prices["apple"] + prices["banana"])
}
