// lesson14 ch01 解答
package main

import "fmt"

func main() {
	// Python の [x * 3 for x in range(1, 11)] に相当します。
	// Go では空スライスを用意し、for で回して append するのが定石です。
	// 行数は増えますが、処理の流れが上から順にそのまま読めます。
	tripled := []int{}
	for i := range 10 {
		tripled = append(tripled, (i+1)*3)
	}
	fmt.Println(tripled)
}
