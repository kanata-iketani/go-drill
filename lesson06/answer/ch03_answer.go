// lesson06 ch03 解答
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ParseFloat も Atoi と同じく (値, error) の 2 値を返します。
	// 第 2 引数の 64 は float64 の精度指定で、Go では常にこう書きます。
	a, err := strconv.ParseFloat("1.5", 64)
	if err != nil {
		fmt.Println("変換に失敗しました")
		return
	}
	b, err := strconv.ParseFloat("2.25", 64)
	if err != nil {
		fmt.Println("変換に失敗しました")
		return
	}
	// float64 の表示は %.2f のように桁数を明示すると出力が安定します。
	fmt.Printf("%.2f\n", a+b)
}
