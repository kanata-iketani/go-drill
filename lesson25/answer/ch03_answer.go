// lesson25 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 引数は値渡しでコピーされるため、int のまま受け取っても呼び出し元は変わりません。
// ポインタ *int を受け取り、*p で指し先そのものを書き換えることで、
// 呼び出し元の変数 x に変更を「届かせて」います。
func addTen(p *int) {
	*p += 10
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	x := n
	fmt.Println("before:", x)
	addTen(&x) // &x で x の場所を渡す
	fmt.Println("after:", x)
}
