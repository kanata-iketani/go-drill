// lesson21 ch03 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Name string
	HP   int
}

// __init__ の代わりはこの「ただの関数」です。HP=100 という初期化ルールを
// ここに集めておけば、作る側が初期値を覚えておく必要がなくなります。
// ポインタを返すので、呼び出し先で HP を減らしても同じ 1 体を指し続けます。
func NewPlayer(name string) *Player {
	return &Player{Name: name, HP: 100}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())
	name := fields[0]
	damage, _ := strconv.Atoi(fields[1])

	p := NewPlayer(name)
	p.HP -= damage
	fmt.Printf("%s HP:%d\n", p.Name, p.HP)
}
