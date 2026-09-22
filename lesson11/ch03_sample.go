// lesson11 ch03: slices.Contains（スライスの in の代わり）
package main

import (
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	// x in fruits の代わりに slices.Contains を使います
	fmt.Println(slices.Contains(fruits, "banana"))
	fmt.Println(slices.Contains(fruits, "grape"))

	// int のスライスでも同じ形です
	nums := []int{10, 20, 30}
	if slices.Contains(nums, 20) {
		fmt.Println("20 が含まれています")
	}
}
