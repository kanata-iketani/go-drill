// lesson24 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()

	// 例外を捕まえるのではなく、戻り値の err を調べます。
	// 「失敗したら何をするか」を呼び出した直後にその場で書くのが
	// if err != nil の型です。err を無視しないことが Go の作法です。
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("数値ではありません")
		return
	}
	fmt.Println(n * 2)
}
