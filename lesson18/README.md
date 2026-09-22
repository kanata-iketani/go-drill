# Lesson 18 関数

Go の関数は引数と戻り値に型を書きます（戻り値なしなら型を書きません）。値を 2 つ以上返す多値返却と、lambda に当たる無名関数も学びます。

## ch01 func と引数の型

Go で関数を定義するには `func` を使います。Python の `def` との一番の差は、**引数に型を書く**ことです。引数名の後ろに型を置き、同じ型が続くときは `a, b int` とまとめて書けます。型の合わない値を渡す呼び出しはコンパイルエラーになるので、実行する前に間違いに気づけます。🔴 関数定義はこの先すべての土台になるので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
def greet(name, times):
    for i in range(times):
        print("Hello,", name)
```

Go ではこう書く:

```go
func greet(name string, times int) {
	for range times {
		fmt.Println("Hello,", name)
	}
}
```

## ch02 戻り値の型を書く

値を返す関数は、引数リストの後ろに**戻り値の型**を書きます。Python は `return` するだけで型を書きませんでしたが、Go は `func add(a, b int) int` のように「何型を返すか」を宣言します。宣言と違う型を返すとコンパイルエラーです。**戻り値がない関数は型を書きません**。`func hello()` のように引数リストで止めます。🔴 「返すなら型を書く・返さないなら書かない」を手に覚えさせましょう。

Python ではこう書いた:

```python
def add(a, b):
    return a + b
```

Go ではこう書く:

```go
func add(a, b int) int {
	return a + b
}

// 戻り値なしなら型を書かない
func hello() {
	fmt.Println("hello")
}
```

## ch03 多値を返す関数

Go の関数は**値を 2 つ以上まとめて返せます**。Python ではタプルで `return q, r` と書きましたが、Go では戻り値の型を `(int, int)` と括弧で並べ、受け取る側は多値代入 `q, r := divmod(a, b)` で受けます。`strconv.Atoi` が `(int, error)` を返していたのは、まさにこの仕組みです。🔴 「結果とエラーを一緒に返す」という Go の基本形につながるので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
def divide(a, b):
    return a // b, a % b

q, r = divide(17, 5)
```

Go ではこう書く:

```go
func divide(a, b int) (int, int) {
	return a / b, a % b
}

q, r := divide(17, 5)
```

## ch04 無名関数（lambda の代わり）

名前を付けずにその場で定義する関数を**無名関数**といいます。Python の `lambda` に相当しますが、lambda が 1 つの式しか書けなかったのに対し、Go の無名関数は**普通の関数と同じ書き方で複数行書けます**。変数に代入すれば `square(3)` のように呼び出せますし、関数の引数として渡すこともできます。🟡 「関数も値の一種として変数に入る」という理屈を理解しましょう。

Python ではこう書いた:

```python
square = lambda x: x * x
print(square(3))
```

Go ではこう書く:

```go
square := func(x int) int {
	return x * x
}
fmt.Println(square(3))
```
