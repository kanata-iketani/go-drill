// lesson16 ch06: キーを集めてソートして出力
package main

import (
	"fmt"
	"slices"
)

func main() {
	scores := map[string]int{"suzuki": 90, "sato": 80, "tanaka": 70}

	// 1. キーだけをスライスに集める（range はキーだけも受け取れます）
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}

	// 2. ソートしてから、その順に map を引く
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println(k, scores[k])
	}
}
