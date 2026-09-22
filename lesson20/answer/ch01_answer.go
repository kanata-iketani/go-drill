// lesson20 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// n はコピーなので、中で n += 10 しても呼び出し元には届きません。
// だから「変えた結果」は return で返し、使う側が受け取ります。
func addTen(n int) int {
	n += 10
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	x, _ := strconv.Atoi(strings.Fields(sc.Text())[0])

	fmt.Println(addTen(x))
	fmt.Println(x) // コピーが渡ったので x は元のまま
}
