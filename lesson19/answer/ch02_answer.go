// lesson19 ch02 解答
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
	x, _ := strconv.Atoi(strings.Fields(sc.Text())[0])

	{
		// 右辺の x は外の x、左辺の x はこのブロックだけの新しい変数です。
		// := が「宣言」だからこそ、同じ名前でも別物になります（シャドーイング）。
		x := x * 10
		fmt.Println(x)
	}
	// ブロックを出ると外の x はそのまま。シャドーイングは外に影響しません
	fmt.Println(x)
}
