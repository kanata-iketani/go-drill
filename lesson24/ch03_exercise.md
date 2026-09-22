# lesson24 ch03 演習: fmt.Errorf("%w") と errors.Is（エラーの包み直しと判定）

1 行目にユーザー名が入力されます。`gopher` のときだけ成功する関数 `find` を用意しました。main で `find` を呼び、err が nil なら `<名前> が見つかりました` と 1 行出力してください。err が nil でなければ、1 行目にエラーメッセージを出力し、さらに `errors.Is(err, ErrNotFound)` が true なら 2 行目に `原因は ErrNotFound です` と出力してください。`%w` で包んでも原因を判定できることを確かめる問題です。

**入力される値**

```text
ユーザー名
```

**入力例1**

```text
gopher
```

**期待出力1**

```text
gopher が見つかりました
```

## テストケース

### ケース 1

入力:

```text
gopher
```

期待出力:

```text
gopher が見つかりました
```

### ケース 2

入力:

```text
alice
```

期待出力:

```text
検索 alice: 見つかりません
原因は ErrNotFound です
```

### ケース 3

入力:

```text
bob
```

期待出力:

```text
検索 bob: 見つかりません
原因は ErrNotFound です
```
