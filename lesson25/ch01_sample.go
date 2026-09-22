// lesson25 ch01: 値型は代入でコピーされる
package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// 構造体は値型: 代入で中身が丸ごとコピーされます
	p := Point{X: 1, Y: 2}
	q := p
	q.X = 100
	fmt.Println(p.X, q.X) // 1 100（p は変わらない）

	// 配列（固定長）も値型です
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a[0], b[0]) // 1 100

	// int も同じ: コピーを変えても元は変わりません
	x := 10
	y := x
	y = 999
	fmt.Println(x, y) // 10 999
}
