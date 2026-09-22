// lesson18 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 引数には名前のあとに型を書きます。同じ int が続くので a, b int とまとめました。
// 型を書いておくと printSum("x", 1) のような誤った呼び出しが
// コンパイル時に弾かれるのが Python の def との大きな差です。
func printSum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])
	printSum(a, b)
}
