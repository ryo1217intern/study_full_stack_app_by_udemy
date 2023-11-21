# 環境構築

1. git の初期化
   `git init`

2. go の初期化
   `go mod init go-rest-api`

今回は version を udemy 講座と合わせて`version:1.20`にしています 3. go tools のインストール
a. goimports
`golang.org/x/tools/cmd/goimports@latest`
goimports はコード本文内で使用しているライブラリを自動で imports してくれる.

b.golint
`go install golang.org/x/lint/golint@latest`
コーディングスタイルの問題や一般的なエラーパターンを検出する。(変数の命名規則、コメントの形式、不要なコード、冗長なコードなど)
