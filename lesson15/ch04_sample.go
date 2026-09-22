// lesson15 ch04: 複数行の入力（N を読んでから N 行読む）
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// 1 行目: 件数 N
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// 続く N 行を読む。Scan() は呼ぶたびに次の行へ進みます
	for i := range n {
		sc.Scan()
		fmt.Printf("%d行目: %s\n", i+1, sc.Text())
	}
}
