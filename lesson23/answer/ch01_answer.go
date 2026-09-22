// lesson23 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

// main 以外の関数は入口にならず、main から呼ばれてはじめて動きます。
// そのため Python のような if __name__ == "__main__": ガードは不要です。
func greet(name string) string {
	return "こんにちは、" + name + " さん"
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()

	// 実行は常に main から始まる、が Go の決まりです
	fmt.Println(greet(name))
}
