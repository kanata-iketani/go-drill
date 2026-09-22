// lesson19 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 関数の外で宣言すると、同じパッケージのどの関数からも見えます。
// Python の global 文に当たるものは Go にはなく、宣言なしで書き換えられます。
var balance = 1000

func apply(n int) {
	balance += n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	for _, f := range fields {
		n, _ := strconv.Atoi(f)
		apply(n)
	}
	fmt.Println(balance)
}
