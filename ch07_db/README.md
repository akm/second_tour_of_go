# データベース

## 事前準備+予習

以下を行っておいてください

1. [Docker Desktop をインストール](https://docs.docker.jp/desktop/install.html)
2. [Get Started - 始めましょう / 概要説明とセットアップ](https://docs.docker.jp/get-started/index.html)
    - `docker/getting-started` を起動してブラウザやcurlでアクセスしてみてください
3. [jq](https://stedolan.github.io/jq/) のインストール

## 概要

この章ではgo言語でデータベースを扱う方法を紹介し、以下のステップで実際にCLIアプリを作ります。

1. [Docker](https://www.docker.com/) を用いたデータベース(MySQL)の開発環境を構築
2. [goose](https://github.com/pressly/goose) を用いたマイグレーション（スキーマ操作等）
3. [Go言語の標準ライブラリのsqlパッケージ](https://pkg.go.dev/database/sql@go1.19.4) を用いたRDBの操作
4. ORMライブラリ[gorm](https://gorm.io/ja_JP/docs/index.html) を用いたRDBの操作

## Dockerでデータベースを構築

まず [dockerhubのmysqlのページ](https://hub.docker.com/_/mysql) を開き、起動方法などが書いてあることを確認します。
ただし、載っているコマンドをそのまま実行しても期待通りに動くとは限りません

1. MySQLサーバーを起動
    `docker run -d --rm --name mysql-server -p 3306:3306 -e MYSQL_DATABASE=testdb1 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:8.0`
1. 起動したサーバーに接続
    - Dockerで起動するMySQLクライアントで起動したサーバーに接続
        `docker run -it --rm mysql mysql testdb1 -u root -h $(docker inspect mysql-server | jq -r '.[].NetworkSettings.IPAddress')`
    - ローカルにインストールされているmysqlクライアントで接続
        `mysql -h 127.0.0.1 -u root testdb1`
1. 操作できることを確かめてみる
    - `SHOW DATABASES;`
    - `SHOW TABLES;`
    - `CREATE TABLE users (id INT, email VARCHAR(255));`
    - `INSERT INTO users (id, email) VALUES (1, 'foo@example.com');`
    - `SELECT * FROM users;`
    - `UPDATE users SET email='bar@example.com' WHERE id = 1;`
    - `DELETE FROM users WHERE id = 1;`
    - `DROP TABLE users;`
