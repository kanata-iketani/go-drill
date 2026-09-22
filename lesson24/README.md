# Lesson 24 エラー処理

Go に例外は存在せず、`try/except` の代わりに関数が `error` を値として返し、`if err != nil` で分岐します。`raise` の代替は `return err` です。

## ch01 Go に例外はない（error は戻り値）

Go に**例外**は存在しません。`try/except` はなく、失敗しうる関数は結果と一緒に `error`（失敗を表す組み込みインターフェース）を**戻り値として**返します。呼び出し側は `if err != nil` で失敗を判定します（`nil` は「エラーなし」を表す値）。`strconv.Atoi` でずっと書いてきた形こそが、Go のエラー処理そのものです。🔴 `if err != nil` は Go で最も頻出する形です。手が覚えるまで書きましょう。

Python ではこう書いた:

```python
try:
    n = int(s)
except ValueError:
    print("数値ではありません")
```

Go ではこう書く:

```go
n, err := strconv.Atoi(s)
if err != nil {
	fmt.Println("数値ではありません")
}
```

## ch02 errors.New と return err（raise の代替は return）

自分でエラーを作るには `errors.New("メッセージ")` を使います。Python の `raise` に相当するのはエラーを投げることではなく **`return err`**、つまりエラーを戻り値として返すことです。失敗しうる関数は `(結果, error)` の 2 値を返す形にし、成功時はエラーの位置に `nil` を返します。🔴 「失敗したら return 0, errors.New(...)、成功したら return 値, nil」の型を手が覚えるまで書きましょう。

Python ではこう書いた:

```python
def divide(a, b):
    if b == 0:
        raise ZeroDivisionError("0 では割れません")
    return a // b
```

Go ではこう書く:

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("0 では割れません")
	}
	return a / b, nil
}
```

## ch03 fmt.Errorf("%w") と errors.Is（エラーの包み直しと判定）

下位のエラーに文脈を足して返すには、`fmt.Errorf` の書式 `%w`（wrap、エラーを包み込むこと）を使います。包んだ後でも `errors.Is(err, 元のエラー)` で「原因はこのエラーか」を判定できます。Python で `raise RuntimeError("検索失敗") from e` と原因を保ったまま投げ直したことに相当します。🟡 「メッセージを足しても原因は失われない」という仕組みを理解しましょう。

Python ではこう書いた:

```python
try:
    find(name)
except KeyError as e:
    raise RuntimeError("検索失敗") from e
```

Go ではこう書く:

```go
if err != nil {
	return fmt.Errorf("検索失敗: %w", err) // 原因を包む
}
// errors.Is(err, ErrNotFound) で原因を判定できる
```

## ch04 独自エラー型（Error() string を実装）

`error` の正体は、`Error() string` というメソッドを 1 つだけ要求するインターフェースです。自作の構造体にこのメソッドを実装すれば**独自エラー型**になり、「どの値で失敗したか」といった追加情報をエラー自体に持たせられます。Python で `class AgeError(Exception)` と例外クラスを定義したことに相当します。🟡 「エラーもただの型」という見方を理解しましょう。

Python ではこう書いた:

```python
class AgeError(Exception):
    def __init__(self, age):
        self.age = age
```

Go ではこう書く:

```go
type AgeError struct{ Age int }

func (e *AgeError) Error() string {
	return fmt.Sprintf("%d歳は成人ではありません", e.Age)
}
```

## ch05 panic と recover（本当に使う場面は少ない）

`panic` はプログラムを即座に停止へ向かわせる仕組み、`recover` は `defer` した関数の中でだけ panic を捕まえて復帰できる関数です。Python の例外と `except` に似ていますが、Go では通常の失敗は `error` で返すのが原則で、panic は「続行できないバグ」のような場面に限られます。本当に使う場面は少ない機能です。⚪ 存在と使いどころだけ知っておきましょう。

Python ではこう書いた:

```python
try:
    risky()
except Exception as e:
    print("回復:", e)
```

Go ではこう書く:

```go
defer func() {
	if r := recover(); r != nil {
		fmt.Println("recover:", r)
	}
}()
```
