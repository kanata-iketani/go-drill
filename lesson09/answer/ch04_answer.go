// lesson09 ch04 解答
package main

import "fmt"

// タプル (3, 5) の代わりに、X と Y という名前を持つ型を定義します。
// 位置番号ではなく名前で読めるので、何の値かが一目で分かります。
type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 3, Y: 5}
	fmt.Printf("X=%d Y=%d\n", p.X, p.Y)
}
