// lesson20 ch05 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 再帰は「基底条件を先に書く」が鉄則です。n == 0 で必ず止まるので、
// fact(5) → 5 * fact(4) → ... → 5 * 4 * 3 * 2 * 1 * fact(0) と展開されて答えが出ます。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(strings.Fields(sc.Text())[0])
	fmt.Println(fact(n))
}
