// lesson07 ch06: 要素の削除（pop・remove・del はない）
package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"a", "b", "c", "d"}

	// pop() に当たる末尾削除は「末尾を除いた切り出し」で書きます
	s = s[:len(s)-1]
	fmt.Println(s)

	// del a[1] に当たる削除は slices.Delete。1 番目から 2-1 番目までを消します
	s = slices.Delete(s, 1, 2)
	fmt.Println(s)
}
