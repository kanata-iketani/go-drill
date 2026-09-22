# Lesson 20 引数と再帰

Go の引数は値渡し（コピー）で、呼び出し元を変えるにはポインタを渡します。キーワード引数・デフォルト引数はなく、オプション構造体で代替します。

## ch01 値渡し（引数はコピーされる）

Go の関数の引数は**値渡し**です。関数を呼ぶと引数の値が**コピー**されて渡り、関数の中でコピーを書き換えても呼び出し元の変数は変わりません。Python でも数値や文字列は同じ振る舞いでしたが、Go では構造体なども丸ごとコピーされます。変更した結果が欲しいときは、戻り値で返して受け取り直すのが基本です。🟡 「関数の中の n は別物」という理屈を理解しましょう。

Python ではこう書いた:

```python
def add_ten(n):
    n += 10  # 呼び出し元の x は変わらない

x = 5
add_ten(x)
print(x)  # 5
```

Go ではこう書く:

```go
func addTen(n int) int {
	n += 10 // 書き換わるのはコピーだけ
	return n
}

x := 5
x = addTen(x) // 戻り値で受け取り直す
```

## ch02 ポインタ渡し（*int と &）

呼び出し元の変数そのものを関数から書き換えたいときは**ポインタ**（変数のメモリ上の場所を指す値）を渡します。`&x` で x のポインタを取り、受け取る引数の型は `*int`、関数の中では `*p` で指した先を読み書きします。Python ではリストなどミュータブルな値を渡して間接的に書き換えましたが、Go では「呼び出し元を変えたい」という意図を `&` と `*` で明示します。🟡 「ポインタを渡す＝場所を教える」という理屈を理解しましょう。

Python ではこう書いた:

```python
def double(nums):
    nums[0] *= 2  # リスト経由でしか書き換えられない

nums = [5]
double(nums)
```

Go ではこう書く:

```go
func double(p *int) {
	*p *= 2 // ポインタの指す先を書き換える
}

x := 5
double(&x) // x の場所を渡す
```

## ch03 可変長引数 ...int（*args の代わり）

引数の個数を決めずに受け取るには、最後の引数の型の前に `...` を付けます（**可変長引数**）。Python の `*args` に相当し、関数の中では普通のスライスとして扱えます。呼び出しは `total(1, 2, 3)` のように好きな個数で書け、スライスをばらして渡すときは Python の `total(*nums)` に当たる `total(nums...)` と書きます。可変長引数は引数リストの最後に 1 つだけ置けます。🟡

Python ではこう書いた:

```python
def total(*args):
    return sum(args)

nums = [1, 2, 3]
print(total(*nums))
```

Go ではこう書く:

```go
func total(nums ...int) int {
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}

fmt.Println(total(nums...)) // スライスをばらして渡す
```

## ch04 キーワード引数・デフォルト引数はない

Python の `greet(name, prefix="Hello")` のような**デフォルト引数**や、`greet(name="太郎")` のような**キーワード引数**は、**Go にはありません**。引数はすべて順番どおりに渡します。省略できる設定が欲しいときは、設定をまとめた**オプション構造体**（詳しくは次のレッスンで学びます）を引数にして、ゼロ値（`""` や 0）だったら関数側でデフォルトを補う、という書き方で代替するのが定番です。🟡

Python ではこう書いた:

```python
def greet(name, prefix="Hello"):
    print(f"{prefix}, {name}!")

greet("Taro")
greet("Hana", prefix="Hi")
```

Go ではこう書く:

```go
type Options struct {
	Prefix string
}

func greet(name string, opt Options) {
	if opt.Prefix == "" {
		opt.Prefix = "Hello" // ゼロ値ならデフォルトを補う
	}
	fmt.Printf("%s, %s!\n", opt.Prefix, name)
}

greet("Taro", Options{})
greet("Hana", Options{Prefix: "Hi"})
```

## ch05 再帰

関数が自分自身を呼び出すことを**再帰**といいます。書き方の考え方は Python と同じで、まず「これ以上自分を呼ばずに答えを返す条件（基底条件）」を書き、それ以外では少し小さくした問題で自分を呼びます。Go では引数と戻り値に型を書く点だけが違います。基底条件を忘れると自分を呼び続けてスタックオーバーフローで落ちるので、必ず先頭に書きましょう。🟡 再帰は理屈が命です。

Python ではこう書いた:

```python
def fact(n):
    if n == 0:
        return 1
    return n * fact(n - 1)
```

Go ではこう書く:

```go
func fact(n int) int {
	if n == 0 {
		return 1 // 基底条件
	}
	return n * fact(n-1)
}
```
