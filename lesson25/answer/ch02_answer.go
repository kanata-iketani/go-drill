// lesson25 ch02 解答
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
	n, _ := strconv.Atoi(strings.Fields(sc.Text())[0])
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a := make([]int, n)
	for i := range n {
		a[i], _ = strconv.Atoi(fields[i])
	}

	// b := a でコピーされるのはヘッダ（配列の場所・長さ・容量）だけです。
	// 中身の配列は a と b で共有されているので、b[0] への書き込みは
	// a からも見えます。「参照的」に振る舞う理由はこの共有にあります。
	b := a
	b[0] *= 10
	fmt.Println(a)
	fmt.Println(b)
}
