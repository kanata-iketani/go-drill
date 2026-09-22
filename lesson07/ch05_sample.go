// lesson07 ch05: append で結合（s2...）
package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := []int{3, 4}

	// b... で b の要素を展開し、1 つずつ追加した扱いになります
	c := append(a, b...)
	fmt.Println(c)
	fmt.Println(len(c))

	// 文字列のスライスでも同じです
	x := []string{"朝", "昼"}
	x = append(x, []string{"夜"}...)
	fmt.Println(x)
}
