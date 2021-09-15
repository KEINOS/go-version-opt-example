# Go Example for Simple Version Option Flag

`go install` でインストールされたバイナリでもバージョン情報を Git のタグに連動させるサンプル。

```shellsession
$ cd /tmp
$ go install github.com/KEINOS/go-version-opt-example@latest
...
$ cd
$ go-version-opt-example --version
go-version-opt-example v0.1.0
```
