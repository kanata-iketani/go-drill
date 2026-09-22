// lesson06 ch02: strconv.Atoi と err（文字列→数値）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi は (int, error) の 2 つの値を返します
	n, err := strconv.Atoi("300")
	if err != nil {
		fmt.Println("変換に失敗しました")
		return
	}
	fmt.Println(n + 100) // 数値になっているので計算できます

	// 数値にできない文字列では err に失敗の理由が入ります
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("abc は数値に変換できません")
	}
}
