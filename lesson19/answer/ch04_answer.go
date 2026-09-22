// lesson19 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	total := 0
	for sc.Scan() {
		// n と err はこの if の中だけの変数です。中で使い切るならこの形が最短で、
		// 外に持ち出したくなったら「先に宣言して = で代入」に切り替えます。
		// うっかり := で外の変数を隠すのが実務でよくある落とし穴です。
		if n, err := strconv.Atoi(sc.Text()); err == nil {
			total += n
		}
	}
	fmt.Println(total)
}
