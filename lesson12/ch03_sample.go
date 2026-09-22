// lesson12 ch03: continue（次の繰り返しへ進む）
package main

import "fmt"

func main() {
	// i を先に更新してから continue するのが安全な形です
	i := 0
	for i < 5 {
		i++
		if i == 3 {
			continue // 3 のときは Println を飛ばして次の繰り返しへ
		}
		fmt.Println(i)
	}

	// もし i++ より前で continue すると、i が増えず無限ループになります
	fmt.Println("おわり")
}
