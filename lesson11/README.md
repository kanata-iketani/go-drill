# Lesson 11 ブール型

`and` `or` `not` は `&&` `||` `!` になります。Go に `in` 演算子はなく、スライスは `slices.Contains`、文字列は `strings.Contains` で調べます。

## ch01 bool 型と比較演算子

真偽を表す型が `bool`（真か偽かの 2 値だけを持つ型）で、値は小文字の `true` と `false` です。Python の `True` / `False` とは先頭の大文字小文字が違います。比較演算子 `==` `!=` `<` `<=` `>` `>=` は Python と同じで、結果は bool 型になります。ただし `0 < x < 10` のような比較の連結はできず、次章で学ぶ `&&` で 2 つの比較に分けます。🔴 比較の結果を bool として変数に入れて使う形を、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
is_adult = True
print(age >= 20)
```

Go ではこう書く:

```go
isAdult := true
fmt.Println(age >= 20)
```

## ch02 && と || と !（and or not の代わり）

Python の `and` `or` `not` は、Go では記号の `&&` `||` `!` になります。`&&` は両方が true のとき true、`||` はどちらかが true のとき true、`!` は真偽を反転します。単語のまま `and` と書くとコンパイルエラーです。前章で触れたとおり比較の連結はできないため、Python の `0 < x < 10` は `x > 0 && x < 10` と書きます。🔴 この 3 つの記号は条件分岐の中核なので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
if x > 0 and x < 10:
    print("1桁")
```

Go ではこう書く:

```go
if x > 0 && x < 10 {
	fmt.Println("1桁")
}
```

## ch03 slices.Contains（スライスの in の代わり）

Go に Python の `in` 演算子はありません。スライスに値が含まれるかは、標準ライブラリの `slices` パッケージにある `slices.Contains(スライス, 値)` で調べます。戻り値は bool なので、そのまま if の条件に使えます。使うには `import "slices"` が必要です。Go 1.21 より前はこの関数がなく、for で 1 件ずつ比較して探していました。🟡 「演算子ではなく関数で調べる」という Go の流儀を理解しましょう。

Python ではこう書いた:

```python
if "banana" in fruits:
    print("あります")
```

Go ではこう書く:

```go
if slices.Contains(fruits, "banana") {
	fmt.Println("あります")
}
```

## ch04 strings.Contains（文字列の in の代わり）

文字列に部分文字列が含まれるかを調べる Python の `"go" in s` は、Go では `strings.Contains(s, "go")` と書きます。前章の `slices.Contains` はスライス用、こちらは文字列用で、`in` の対象の型によって使う関数が変わります。`strings` パッケージには、先頭が一致するかを調べる `HasPrefix`、末尾を調べる `HasSuffix` もあります。🟡 「in の代わりは型ごとに関数を選ぶ」という考え方を理解しましょう。

Python ではこう書いた:

```python
if "go" in s:
    print("含まれます")
```

Go ではこう書く:

```go
if strings.Contains(s, "go") {
	fmt.Println("含まれます")
}
```
