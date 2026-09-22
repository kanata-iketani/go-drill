// lesson20 ch05: 再帰
package main

import "fmt"

// fact は n の階乗（n * (n-1) * ... * 1）を再帰で計算します
func fact(n int) int {
	if n == 0 {
		return 1 // 基底条件: これがないと無限に呼び続ける
	}
	return n * fact(n-1) // 少し小さい問題を自分に任せる
}

// カウントダウンも再帰で書けます
func countdown(n int) {
	if n < 0 {
		return
	}
	fmt.Println(n)
	countdown(n - 1)
}

func main() {
	fmt.Println(fact(5)) // 120
	countdown(3)
}
