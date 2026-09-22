// lesson23 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

// init は main より先に自動実行されます。呼び出しコードは書きません。
// Python でモジュールのトップレベルに書いた処理が import 時に走るのと
// 同じ役割を、Go では init という関数の形で担います。
func init() {
	fmt.Println("設定を読み込みました")
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()
	fmt.Println("ようこそ、" + name + " さん")
}
