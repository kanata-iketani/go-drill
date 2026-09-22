// lesson21 ch02: メソッドとレシーバ（self の代わり）
package main

import "fmt"

type Rect struct {
	W, H int
}

// 値レシーバ: 読むだけのメソッド。r が self の代わりです
func (r Rect) Area() int {
	return r.W * r.H
}

// ポインタレシーバ: フィールドを書き換えるメソッドは *Rect にします
func (r *Rect) Scale(n int) {
	r.W *= n
	r.H *= n
}

func main() {
	r := Rect{W: 2, H: 3}
	fmt.Println(r.Area()) // 6

	r.Scale(2)            // 呼び出し側の書き方は値レシーバと同じ
	fmt.Println(r.Area()) // 24
}
