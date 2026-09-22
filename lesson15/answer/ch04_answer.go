// lesson15 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// 1 行目の N を先に読み、その回数だけ Scan を繰り返すのが定型です。
	// Scan() が呼ぶたびに次の行へ進むので、ループの中で読み進められます。
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sum := 0
	for range n {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		sum += v
	}
	fmt.Println(sum)
}
