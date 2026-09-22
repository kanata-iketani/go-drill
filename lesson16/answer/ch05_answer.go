// lesson16 ch05 解答
package main

import (
	"bufio"
	"fmt"
	"os"
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

	// range の順序は不定ですが、合計は足す順番に関係なく同じ結果になります。
	// このように「順序に依存しない集計」だけが range の安全な使いどころです。
	sum := 0
	for _, v := range scores {
		sum += v
	}
	fmt.Println(sum)
}
