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

## :exclamation: Hello, World!

以下の内容でmain.goを作成し、実行してください。また実行形式のファイルを作成した上でそれを実行してください。

```golang
package main

func main() {
  println("Hello, World!")
}
```

### 使用するコマンド

- ビルド(=実行形式ファイルを作成)
  - `go build .`
- 実行形式ファイルを作成せずに実行
  - `go run .`

### 実行形式のファイルの実行

環境 | コマンド
---------|--------
Linux/Mac | `./second_tour_of_go` 
Windows   | `./second_tour_of_go.exe`


## :question: Hello, Arguments!

以下の条件を満たすように変更してください(実行形式のファイルを ch02_first_cli_app とします)。

```
./ch02_first_cli_app Golang # => Hello, Golang!
./ch02_first_cli_app        # => Hello, someone!
```

### ヒント

[標準ライブラリ](https://pkg.go.dev/std) の `os` パッケージの変数 `Args` を使います。
また `fmt` パッケージを使うとよりシンプルに書けると思います。


## :question: 足し算アプリ

以下の条件を満たすように変更してください(実行形式のファイルを ch02_first_cli_app とします)。

```
./ch02_first_cli_app 1 2 # => Result: 3
./ch02_first_cli_app 1 a # => ERROR: "a" is not a number
./ch02_first_cli_app x a # => ERROR: "x" is not a numeric
./ch02_first_cli_app x   # => (ヘルプを表示)
./ch02_first_cli_app     # => (ヘルプを表示)
```

ERRORやヘルプを表示する際の終了コードは1（それ以外は0）とする

### ヘルプの内容

```
USAGE:
./ch02_first_cli X Y
   Shows sum of X and Y
   X and Y must be number
```

### ヒント

以下の関数を利用

- `strconv.Atoi`
- `os.Exit`
- `fmt.Printf`

## :question: 足し算と引き算のサポート

実行形式のファイルを ch02_first_cli_app とする

```
./ch02_first_cli_app add 1 2      # => addition: 3
./ch02_first_cli_app subtract 1 2 # => subtraction: -1
./ch02_first_cli_app foo 1 2      # => (ヘルプを表示)
./ch02_first_cli_app add 1 a      # => ERROR: "a" is not a number
./ch02_first_cli_app add x a      # => ERROR: "x" is not a number
./ch02_first_cli_app add x        # => (ヘルプを表示)
./ch02_first_cli_app add          # => (ヘルプを表示)
./ch02_first_cli_app subtract 1 a # => ERROR: "a" is not a number
./ch02_first_cli_app subtract x a # => ERROR: "x" is not a number
./ch02_first_cli_app subtract x    # => (ヘルプを表示)
./ch02_first_cli_app subtract      # => (ヘルプを表示)
```

ERRORやヘルプを表示する際の終了コードは1とする

### ヘルプの内容

```
USAGE:
./ch02_first_cli_app (add|subtract) X Y
   Shows addition or subtraction with X and Y
   X and Y must be number
```

## :question: 条件分岐を減らす

この時点での実装では、 main関数でswitch文を２つ使うような実装になっていると思われます。
これをよりシンプルにする方法を考えてください。ただし、出力結果の条件として以下を追加します。

```
./ch02_first_cli_app foo 1 a      # => (ヘルプを表示)
./ch02_first_cli_app foo x 2      # => (ヘルプを表示)
```

### ヒント

インターフェイスを使ってください。


## :question: mapとif文を使ってswitch文を減らす

```golang
  var calc Calculation
  switch os.Args[1] {
  case "add":
    calc = &Addition{}
  case "subtract":
    calc = &Subtraction{}
  default:
    showHelp()
    os.Exit(1)
  }
```

となっている部分を map と if文を使って書き直してください。

## :question: calc.go についてのテスト calc_test.go を追加して、テストを書いてください

calc.go で実装されている機能についてテストを書いてみてください。
もし必要があればcalc.goを変更しても構いません。


## :question: サブコマンド multiply を実装してください

以下の条件を満たすように変更してください(実行形式のファイルを ch02_first_cli_app とします)。
また、テストも作成してください。

```
./ch02_first_cli_app add 1 2      # => addition: 3
./ch02_first_cli_app subtract 1 2 # => subtraction: -1
./ch02_first_cli_app multiply 2 3 # => multiplication: 6
./ch02_first_cli_app foo 1 2      # => (ヘルプを表示)
./ch02_first_cli_app foo 1 a      # => (ヘルプを表示)
./ch02_first_cli_app foo x 2      # => (ヘルプを表示)
./ch02_first_cli_app add 1 a      # => ERROR: "a" is not a number
./ch02_first_cli_app add x a      # => ERROR: "x" is not a number
./ch02_first_cli_app add x        # => (ヘルプを表示)
./ch02_first_cli_app add          # => (ヘルプを表示)
./ch02_first_cli_app subtract 1 a # => ERROR: "a" is not a number
./ch02_first_cli_app subtract x a # => ERROR: "x" is not a number
./ch02_first_cli_app subtract x    # => (ヘルプを表示)
./ch02_first_cli_app subtract      # => (ヘルプを表示)
./ch02_first_cli_app multiply 1 a # => ERROR: "a" is not a number
./ch02_first_cli_app multiply x a # => ERROR: "x" is not a number
./ch02_first_cli_app multiply x    # => (ヘルプを表示)
./ch02_first_cli_app multiply      # => (ヘルプを表示)
```

ERRORやヘルプを表示する際の終了コードは1とします。

### ヘルプの内容

```
USAGE:
./ch02_first_cli_app (add|subtract|multiply) X Y
   Shows the result of calculation with X and Y
   X and Y must be number
```
