// lesson17 ch02 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	words := strings.Fields(sc.Text())

	// seen が set の代用です。初めて見た単語だけ unique に追加することで、
	// Python の set(lst) と違い、元の並び順を保ったまま重複を除去できます。
	seen := map[string]bool{}
	unique := []string{}
	for _, w := range words {
		if !seen[w] {
			seen[w] = true
			unique = append(unique, w)
		}
	}
	fmt.Println(strings.Join(unique, " "))
}
