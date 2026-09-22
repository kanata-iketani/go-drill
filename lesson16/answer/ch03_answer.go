// lesson16 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	menu := map[string]int{"coffee": 480, "tea": 400, "cocoa": 520}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	item := sc.Text()

	// 存在しないキーは 0 が返るだけでエラーにならないため、
	// 「メニューにない」ことを知るにはカンマ ok イディオムが必須です。
	if v, ok := menu[item]; ok {
		fmt.Printf("%d円\n", v)
	} else {
		fmt.Println("ありません")
	}
}
