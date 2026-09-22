// lesson16 ch06 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	scores := make(map[string]int)
	for range 3 {
		sc.Scan()
		f := strings.Fields(sc.Text())
		v, _ := strconv.Atoi(f[1])
		scores[f[0]] = v
	}

	// map の range 順は不定なので、そのまま出力すると結果が毎回変わりえます。
	// 「キーを集める → ソート → その順に引く」が順序を固定する定石です。
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println(k, scores[k])
	}
}
