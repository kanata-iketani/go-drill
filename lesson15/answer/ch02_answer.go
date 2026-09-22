// lesson15 ch02 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	// Text() のままでは文字列なので +1 できません。Atoi で int にします。
	// 入力形式が保証されているスキルチェックでは err は _ で捨てるのが定石です。
	n, _ := strconv.Atoi(sc.Text())

	fmt.Printf("来年は%d歳です\n", n+1)
}
