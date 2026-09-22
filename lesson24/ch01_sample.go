// lesson24 ch01: 例外はなく、error は戻り値
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Atoi は (int, error) の 2 値を返します
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("変換失敗:", err)
	} else {
		fmt.Println("変換成功:", n)
	}

	// 失敗すると err に nil 以外の値が入ります
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("abc は数値にできません")
	}
}
