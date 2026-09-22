// lesson21 ch03: コンストラクタ関数 NewX
package main

import "fmt"

type Player struct {
	Name string
	HP   int
}

// New + 型名 がコンストラクタ関数の慣習です。
// HP の初期値 100 のような「決まりごと」をここに集めます
func NewPlayer(name string) *Player {
	return &Player{Name: name, HP: 100}
}

func main() {
	p := NewPlayer("勇者")
	fmt.Println(p.Name, p.HP) // 勇者 100

	q := NewPlayer("魔法使い")
	q.HP -= 30
	fmt.Println(q.Name, q.HP) // 魔法使い 70
}
