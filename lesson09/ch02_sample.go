// lesson09 ch02: 多値代入とアンパック
package main

import "fmt"

func minmax() (int, int) {
	return 1, 9
}

func main() {
	// アンパックに当たる多値代入。個数が合わないとコンパイルエラーです
	low, high := minmax()
	fmt.Println(low, high)

	// いらない値は _ で捨てます
	_, top := minmax()
	fmt.Println(top)

	// 値の入れ替えは Python と同じ書き方です
	a, b := 10, 20
	a, b = b, a
	fmt.Println(a, b)
}
