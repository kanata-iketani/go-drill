// lesson24 ch05 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(n int) int {
	if n < 0 {
		panic("負の数は扱えません")
	}
	return n * 10
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// recover は defer した関数の中でしか効きません。
	// panic が起きると main の残りは実行されず、この defer だけが走ります。
	// 通常の失敗は error で返すのが原則で、panic はバグ相当の場面だけです。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()

	fmt.Println(check(n))
	fmt.Println("正常に終了しました")
}
