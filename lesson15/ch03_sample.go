// lesson15 ch03: 1 行に複数の値（strings.Fields）
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

	// "3 5" のような 1 行を空白で分割して []string にします
	f := strings.Fields(sc.Text())
	a, _ := strconv.Atoi(f[0])
	b, _ := strconv.Atoi(f[1])

	fmt.Println("個数:", len(f))
	fmt.Println("和:", a+b)
}
