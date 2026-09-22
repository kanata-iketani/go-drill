// lesson18 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 商と余りは「セットで欲しい」結果なので、2 回に分けず一度に返します。
// 戻り値の型 (int, int) と return a / b, a % b の個数は必ず一致させます。
func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])
	// 多値代入で受け取ります。Atoi の (int, error) と同じ形です
	q, r := divide(a, b)
	fmt.Println(q, r)
}
