# 第1章 初めてのテスト

## セットアップ

1. https://go.dev/doc/install の指示に従って Go 1.19 をインストール
    [asdf](https://asdf-vm.com/)の[golangプラグイン](https://github.com/kennyp/asdf-golang) あるいは [goenv](https://github.com/syndbg/goenv) などを使うことを推奨
1. `go` コマンドのバージョンを確認
    1. `ターミナル` (Mac/Linux) or `コマンドプロンプト` (Windows) を新たに開く
    1. `go version` を実行
    1. インストールしたGoのバージョンが表示されることを確認
1. https://github.com/akm/second_tour_of_go/releases/tag/ch01_first_test_01 からソースコードをダウンロード
1. ダウンロードしたコードを解凍
1. ターミナル上でディレクトリ `ch01_first_test` に移動
    `cd path/to/second_tour_of_go-ch01_first_test_01/ch01_first_test`
    (`path/to` は解凍して作られたディレクトリに置き換えて実行してください )
1. `go test` を実行し、以下のような出力が表示されることを確認
    ```
    PASS
    ok      github.com/akm/second_tour_of_go/ch01_first_test        0.341s
    ```
