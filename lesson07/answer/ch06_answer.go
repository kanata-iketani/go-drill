// lesson07 ch06 解答
package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{10, 20, 30, 40}

	// del a[1] に当たる操作。引数 (1, 2) は「1 番目から 2-1 番目まで」の
	// 範囲指定で、戻り値を再代入しないと削除が反映されません。
	nums = slices.Delete(nums, 1, 2)
	fmt.Println(nums)

	// a.pop() に当たる末尾削除は、末尾を除くスライス式で書きます。
	nums = nums[:len(nums)-1]
	fmt.Println(nums)
}
