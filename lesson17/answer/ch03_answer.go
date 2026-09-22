// lesson17 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a := strings.Fields(sc.Text())
	sc.Scan()
	b := strings.Fields(sc.Text())

	// B を set 化しておけば、A を 1 回走査するだけで
	// 「B にある → 積」「B にない → 差」に振り分けられます。
	setB := map[string]bool{}
	for _, w := range b {
		setB[w] = true
	}
	inter := []string{}
	diff := []string{}
	for _, w := range a {
		if setB[w] {
			inter = append(inter, w)
		} else {
			diff = append(diff, w)
		}
	}

	// map を経由した結果は順序が不定になりうるので、ソートして固定します
	slices.Sort(inter)
	slices.Sort(diff)
	fmt.Println(strings.Join(inter, " "))
	fmt.Println(strings.Join(diff, " "))
}
