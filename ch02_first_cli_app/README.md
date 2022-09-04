# 第２章 初めてのCLIアプリ

## CLIとは

Command Line Interfaceの略で、GUIではなくターミナルやPowerShell(あるいはコマンドプロンプト)を介してコマンドを実行してコンピュータを操作する方式を指します。
HTTPサーバーを始め、サーバー上で動作するアプリケーションはGUIを持たないCLIアプリであると言えます。
積極的にターミナルやPowerShellを使ってなれていきましょう。

ただし実際にWebアプリケーションのサーバーを動かす環境としてWindowsを使うことはまずありません。
VMやDockerを使ってLinuxを動かしたり、直接LinuxあるいはMacを使うことを強くお勧めします。

### 参考

- [初心者のためのWindowsコマンドプロンプト使い方入門](https://proengineer.internous.co.jp/content/columnfeature/4962)
- [基本コマンドプロンプト25選！逆引きコマンド一覧](https://proengineer.internous.co.jp/content/columnfeature/5007)

## :exclamation: セットアップ

1. ターミナルあるいはコマンドプロンプトで `ch02_first_cli_app` に移動
2. `go mod init github.com/akm/second_tour_of_go/ch02_first_cli_app` を実行
    - `github.com/akm` の部分は他の文字列に変更してもOKです。
        - Github等で管理するのであれば、自分の環境に合わせて変更してください
