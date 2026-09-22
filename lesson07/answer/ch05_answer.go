// lesson07 ch05 解答
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5}

	// Python の a + b に当たる結合。append は要素を受け取る関数なので、
	// b... と展開して「b の要素を 1 つずつ追加する」形にします。
	c := append(a, b...)
	fmt.Println(c)
}
