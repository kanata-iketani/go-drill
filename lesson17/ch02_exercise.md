# lesson17 ch02 演習: 存在判定と重複除去

標準入力の 1 行に単語が空白区切りで与えられます。重複を取り除き、最初に登場した順序を保ったまま、空白区切りの 1 行で出力してください。

**入力される値**

```text
単語1 単語2 ... 単語N
```

**入力例1**

```text
apple banana apple orange banana
```

**期待出力1**

```text
apple banana orange
```

## テストケース

### ケース 1

入力:

```text
apple banana apple orange banana
```

期待出力:

```text
apple banana orange
```

### ケース 2

入力:

```text
go go go
```

期待出力:

```text
go
```

### ケース 3

入力:

```text
a b c b a d
```

期待出力:

```text
a b c d
```
