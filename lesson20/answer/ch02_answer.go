// lesson20 ch02 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 戻り値ではなく「呼び出し元の x そのもの」を変えたいので、ポインタで受けます。
// *p *= 2 の * は「ポインタの指す先」の意味で、これで元の変数が書き換わります。
func double(p *int) {
	*p *= 2
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	x, _ := strconv.Atoi(strings.Fields(sc.Text())[0])

	double(&x) // & で x の場所を渡す
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}
