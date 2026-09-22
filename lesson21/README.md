# Lesson 21 構造体とメソッド

Go に `class` と `self` はなく、`type X struct` とレシーバ付きメソッドで書きます。`__init__` の代わりは慣習の `New` 関数、公開制御は名前の大文字小文字です。

## ch01 type X struct（class の代わり）

関連するデータをひとまとめにする入れ物が**構造体**です。Python では `class` で作りましたが、**Go に `class` はありません**。`type Person struct { ... }` で**フィールド**（データの項目）を型付きで並べ、`Person{Name: "佐藤", Age: 28}` で初期化し、`p.Name` で読み書きします。`__init__` を書かなくても、この初期化構文でそのまま値を詰められます。🔴 構造体は Go プログラムの中心なので、定義と初期化を手が覚えるまで書きましょう。

Python ではこう書いた:

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

p = Person("佐藤", 28)
print(p.name)
```

Go ではこう書く:

```go
type Person struct {
	Name string
	Age  int
}

p := Person{Name: "佐藤", Age: 28}
fmt.Println(p.Name)
```

## ch02 メソッドとレシーバ（self の代わり）

構造体に結び付いた関数を**メソッド**といい、`func` と関数名の間に**レシーバ**（メソッドを呼ばれる本人）を書きます。Python の `self` に当たるのがレシーバで、**Go に `self` というキーワードはありません**。名前は自分で決めます。フィールドを書き換えるメソッドはポインタレシーバ `(r *Rect)`、読むだけなら値レシーバ `(r Rect)` にします。🔴 `p.Method()` の形は毎日書くので、手が覚えるまで練習しましょう。

Python ではこう書いた:

```python
class Rect:
    def __init__(self, w, h):
        self.w = w
        self.h = h

    def area(self):
        return self.w * self.h
```

Go ではこう書く:

```go
type Rect struct {
	W, H int
}

// (r Rect) がレシーバ。中では r が self の代わり
func (r Rect) Area() int {
	return r.W * r.H
}
```

## ch03 コンストラクタ関数 NewX

Go の構造体には Python の `__init__` のような特別な初期化メソッドは**ありません**。代わりに、初期値を詰めた構造体（へのポインタ）を返す普通の関数を、**`New` で始まる名前**で用意するのが慣習です。これを**コンストラクタ関数**と呼び、デフォルト値の設定や値の検証をこの 1 か所に集められます。戻り値は `*Player` のようにポインタにするのが定番です。🟡 「特別な構文ではなく、ただの関数と命名の約束」という点を理解しましょう。

Python ではこう書いた:

```python
class Player:
    def __init__(self, name):
        self.name = name
        self.hp = 100

p = Player("勇者")
```

Go ではこう書く:

```go
func NewPlayer(name string) *Player {
	return &Player{Name: name, HP: 100}
}

p := NewPlayer("勇者")
```

## ch04 大文字小文字で公開制御

Go では名前の**先頭が大文字か小文字か**で公開範囲が決まります。先頭が大文字（`Name` や `Deposit`）なら他のパッケージからも使える**公開**、小文字（`balance`）ならそのパッケージの中だけの**非公開**です。Python の `_name` は「触らないで」という慣習にすぎませんでしたが、Go では言語仕様として強制されます。フィールドを小文字にして、公開メソッド経由でだけ操作させるのが定番の設計です。🟡 命名がそのまま設計になる、という理屈を押さえましょう。

Python ではこう書いた:

```python
class Wallet:
    def __init__(self):
        self._balance = 0  # _ は「触るな」の目印（強制力なし）

    def deposit(self, n):
        self._balance += n
```

Go ではこう書く:

```go
type Wallet struct {
	balance int // 小文字: パッケージの外から見えない
}

// 大文字: 公開メソッド。外からはこれ経由でだけ触れる
func (w *Wallet) Deposit(n int) {
	w.balance += n
}
```
