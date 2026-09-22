// lesson24 ch02 解答
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// raise の代わりに return でエラーを返します。
// 失敗しうる関数を (結果, error) の 2 値にするのが Go の定番の形で、
// 成功時はエラーの位置に nil を返します。
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("0 では割れません")
	}
	return a / b, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])

	result, err := divide(a, b)
	if err != nil {
		fmt.Println(err) // error は Println でそのままメッセージが出ます
		return
	}
	fmt.Println(result)
}
