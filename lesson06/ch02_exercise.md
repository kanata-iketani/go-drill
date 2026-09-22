# lesson06 ch02 演習: strconv.Atoi と err（文字列→数値）

文字列 `"120"` と `"80"` をそれぞれ `strconv.Atoi` で int に変換し、2 つの合計を出力してください。変換のたびに `if err != nil` のチェックを書くこと。`strconv` を import に追加してください。

**期待出力**

```text
200
```
