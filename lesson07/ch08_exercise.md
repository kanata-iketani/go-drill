# lesson07 ch08 演習: コピーしてからソート（sorted の代替）

スライス `[]string{"c", "a", "b"}` を作り、`slices.Clone` でコピーしたものを `slices.Sort` で並べ替えて 1 行目に出力し、2 行目に元のスライスをそのまま出力して、元が変わっていないことを確認してください。`slices` を import に追加してください。

**期待出力**

```text
[a b c]
[c a b]
```
