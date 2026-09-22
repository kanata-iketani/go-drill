// lesson09 ch04: 構造体で「変更できない組」を表す
package main

import "fmt"

// タプル (x, y) の代わりに、名前付きの型を定義します
type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 3, Y: 5}

	// p[0] ではなく名前でアクセスします
	fmt.Println(p.X)
	fmt.Println(p.Y)

	// 値型なので、コピーを書き換えても元は変わりません
	q := p
	q.X = 100
	fmt.Println(p.X)
}
