// lesson07 ch08 解答
package main

import (
	"fmt"
	"slices"
)

func main() {
	original := []string{"c", "a", "b"}

	// sorted(original) の代わり。Sort はその場で書き換えるため、
	// 元を残したいときは先に Clone で複製してから複製をソートします。
	sorted := slices.Clone(original)
	slices.Sort(sorted)

	fmt.Println(sorted)
	fmt.Println(original)
}
