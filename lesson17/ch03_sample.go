// lesson17 ch03: 和・積・差は自分で書く
package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []string{"apple", "banana", "orange"}
	setB := map[string]bool{"banana": true, "grape": true}

	// 積（A にも B にもある）と差（A にだけある）は 1 回の走査で作れます
	inter := []string{}
	diff := []string{}
	for _, w := range a {
		if setB[w] {
			inter = append(inter, w)
		} else {
			diff = append(diff, w)
		}
	}

	// map を使った処理の出力はソートして順序を固定します
	slices.Sort(inter)
	slices.Sort(diff)
	fmt.Println("積:", inter)
	fmt.Println("差:", diff)
}
