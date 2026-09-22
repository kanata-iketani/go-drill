// lesson06 ch04: strconv.FormatFloat（float64→文字列）
package main

import (
	"fmt"
	"strconv"
)

func main() {
	pi := 3.14159

	// 引数は (値, 表記, 小数桁数, 精度)。'f' は 3.14 のような普通の表記です
	s := strconv.FormatFloat(pi, 'f', 2, 64)
	fmt.Println("円周率はおよそ " + s)

	// 桁数に -1 を渡すと「元の値を復元できる最短の表記」になります
	fmt.Println(strconv.FormatFloat(pi, 'f', -1, 64))
}
