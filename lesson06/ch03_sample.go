// lesson06 ch03: strconv.ParseFloat（文字列→float64）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 第 2 引数の 64 は「float64 の精度で読む」という指定です
	x, err := strconv.ParseFloat("2.5", 64)
	if err != nil {
		fmt.Println("変換に失敗しました")
		return
	}

	// float64 になっているので計算できます
	fmt.Printf("%.1f\n", x+1.5)
}
