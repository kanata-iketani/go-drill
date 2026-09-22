// lesson15 ch05 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// テンプレート: Scanner 生成 + Buffer 拡大までを毎回そのまま書きます
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	f := strings.Fields(sc.Text())

	// 最大値は「最初の要素で初期化」すると負の数だけの入力でも正しく動きます
	sum := 0
	best := 0
	for i := range n {
		v, _ := strconv.Atoi(f[i])
		sum += v
		if i == 0 || v > best {
			best = v
		}
	}
	fmt.Println(sum)
	fmt.Println(best)
}
