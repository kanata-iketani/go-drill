// lesson07 ch07: slices.Sort で並べ替え
package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{3, 1, 4, 1, 5}

	// その場で昇順に並べ替えます。戻り値はありません
	slices.Sort(nums)
	fmt.Println(nums)

	// 文字列も同じ書き方で並べ替えられます
	names := []string{"banana", "apple", "cherry"}
	slices.Sort(names)
	fmt.Println(names)
}
