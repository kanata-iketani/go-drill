// lesson18 ch03: 多値を返す関数
package main

import "fmt"

// 戻り値が複数あるときは型を (int, int) と括弧で並べます
func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	// 受け取る側は多値代入。順番どおりに入ります
	q, r := divide(17, 5)
	fmt.Println("商:", q)
	fmt.Println("余り:", r)

	// 片方だけ欲しいときは _ で捨てます
	q2, _ := divide(100, 7)
	fmt.Println("100÷7 の商:", q2)
}
