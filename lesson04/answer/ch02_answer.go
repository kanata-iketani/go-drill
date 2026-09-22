// lesson04 ch02 解答
package main

import "fmt"

func main() {
	// 関数の中では := が最短の宣言方法で、実務でも主流の書き方です。
	// 3 つとも右辺から型が推論されます（string, int, int）。
	item := "りんご"
	count := 3
	price := 150

	fmt.Println(item, count, "個で", count*price, "円")
}
