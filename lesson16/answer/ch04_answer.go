// lesson16 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stock := map[string]int{"pen": 120, "note": 180, "clip": 50}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()

	// Go の delete は存在しないキーでも安全（何も起きない）なので、
	// Python の del と違い、存在チェックなしでそのまま呼べます。
	delete(stock, name)
	fmt.Println(len(stock))
}
