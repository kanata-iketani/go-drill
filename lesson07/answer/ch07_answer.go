// lesson07 ch07 解答
package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{5, 2, 8, 1}

	// a.sort() と同じ「その場で並べ替え」なので、戻り値を受け取る必要は
	// ありません。並べ替わった nums をそのまま出力します。
	slices.Sort(nums)
	fmt.Println(nums)
}
