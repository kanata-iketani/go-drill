// lesson11 ch03 解答
package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{2, 4, 6, 8}

	// Python の 6 in nums の代わりです。slices.Contains は bool を
	// 返すので、そのまま Println に渡せば true / false が出力されます。
	fmt.Println(slices.Contains(nums, 6))
	fmt.Println(slices.Contains(nums, 5))
}
