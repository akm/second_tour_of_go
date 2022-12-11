# データベース

## 概要

この章ではgo言語でデータベースを扱う方法を紹介し、以下のステップで実際にCLIアプリを作ります。

1. [Docker](https://www.docker.com/) を用いたデータベース(MySQL)の開発環境を構築
2. [goose](https://github.com/pressly/goose) を用いたマイグレーション（スキーマ操作等）
3. [Go言語の標準ライブラリのsqlパッケージ](https://pkg.go.dev/database/sql@go1.19.4) を用いたRDBの操作
4. ORMライブラリ[gorm](https://gorm.io/ja_JP/docs/index.html) を用いたRDBの操作
