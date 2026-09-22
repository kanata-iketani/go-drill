// lesson15 ch03 解答
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
	sc.Scan()

	// 1 行を Fields で []string に割り、必要な要素を Atoi で数値化します。
	// 「Fields で割って Atoi」が複数値入力の定型です。
	f := strings.Fields(sc.Text())
	a, _ := strconv.Atoi(f[0])
	b, _ := strconv.Atoi(f[1])

	fmt.Println(a + b)
	fmt.Println(a * b)
}
