// lesson07 ch08: コピーしてからソート（sorted の代替）
package main

import (
	"fmt"
	"slices"
)

func main() {
	scores := []int{30, 10, 20}

	// sorted(scores) の代わり: 複製を作って、複製の方を並べ替えます
	ranked := slices.Clone(scores)
	slices.Sort(ranked)

	fmt.Println(ranked)
	fmt.Println(scores) // 元の並びは残っています
}
