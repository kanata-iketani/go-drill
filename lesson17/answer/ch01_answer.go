// lesson17 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	members := map[string]bool{"sato": true, "suzuki": true, "tanaka": true}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()

	// map[T]bool の集合は、ないキーがゼロ値 false になるため、
	// Python の in に相当する所属判定が members[name] だけで書けます。
	if members[name] {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
