# lesson15 ch01 演習: bufio.Scanner で 1 行読む

標準入力から名前が 1 行で与えられます。`bufio.Scanner` で読み取り、`こんにちは、{名前}さん` の形式で出力してください。

**入力される値**

```text
名前
```

**入力例1**

```text
Tanaka
```

**期待出力1**

```text
こんにちは、Tanakaさん
```

## テストケース

### ケース 1

入力:

```text
Tanaka
```

期待出力:

```text
こんにちは、Tanakaさん
```

### ケース 2

入力:

```text
Go
```

期待出力:

```text
こんにちは、Goさん
```

### ケース 3

入力:

```text
gopher
```

期待出力:

```text
こんにちは、gopherさん
```
