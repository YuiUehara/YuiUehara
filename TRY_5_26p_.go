package main

// 5章 コマンドラインツール 26p 【TRY】 catコマンドを作ろう

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    //"strings"

    //"strings"//文字列周りの処理
    "path/filepath" //ファイルパスを貰うため
)

// 1.引数でファイルパスの一覧を貰う
    // コマンドライン（コンソールで入力する$ mycat -n hoge~）から文字列を受け取って 文字列を受け取る方法
    //  catコマンドはよくファイルの閲覧に使われるコマンド
    // 
    // 受け取った物を変数nに入れる 変数に入れる方法
    // ルート1 受け取った文字列をファイルパスに変換する 文字列を変換する方法

// p10より参照
    // var msg = flag.String("msg", filepath(), "ここに何かしらの説明")
    // var msg = flag.String("msg", "filepath()", "ここに何かしらの説明")

// var n int //nは整数
var (
    number bool // ?
//  n string //nは文字列
    strCount int
	files []string
)

func init() { //初期化、（プログラム引数の取得する）
    flag.BoolVar(&number, "number", number, "通し番号を付与する")
    flag.IntVar(&strCount, "Count", strCount, "各ファイルにおける出力する文字数")
    flag.Parse() // 解析、Parse()を読んだのでflagのArgs()でフラッグ以外を取得できるそうです。
    files = flag.Args() // フラッグ以外のコマンドライン引数を返す？
}

func main() {
    // 実行ファイルのパスを取得
    path, err := os.Executable() //自分自身の取得
    // ここから
    if err != nil { // エラーが出た時の処理
		log.Fatal(err)
	//	return
    }
    // ここまで何のための物かわからん
    // path = filepath.Dir(path) // 実行ファイルのディレクトリを取得



    // p10より参照(flagパッケージ)
    //flag.Parse()
    //fmt.Println(strings.Repeat(*msg, n))
    ///fmt.Println(flag.Args()
    //n := (ここに受け取った物をいれる）（filepathが入るように)


    for i, file := range files {
        // 行番号、ファイル名を表示させる
        if number {
            fmt.Printf("%v: %v\n", i, file)
        } else {
            fmt.Println(file)
        }
    }

    // 読み込み
    path, err := os.Open(filepath.Join(path, file)) //ディレクトリにファイル名を結合
    //path err := os.Open(filepath.Join("dir", "main.go")) //ディレクトリにファイル名を結合
    if err != nil {
        fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
    }

    // p23より参照（filepathパッケージ）
    // パスを結合する
    //path := filepath.Join("dir", "main.go")
    ///まだ理解していないが、これがないと下記のファイル名の取得ができない。

    //path = filepath.Dir(path) // 実行ファイルのディレクトリを取得
    //fmt.Println(path)

    // ファイル名を取得
    //fmt.Println(filepath.Base(path))


// ここから2.ファイルを順に標準出力に出力させる

    // 出力
    scanner := bufio.NewScanner(path)
/// エラーメッセージ
/// ./main.go:92:29: cannot use path (type string) as type io.Reader in argument to bufio.NewScanner:
/// bufio.NewScannerへの引数でタイプio.Readerとしてパス（タイプ文字列）を使用できません：
/// string does not implement io.Reader (missing Read method)
/// 文字列はio.Readerを実装していません（Readメソッドがありません）

    // TODO: 変数strCountによる出力文字数制限
    for scanner.Scan(){
        fmt.Println(scanner.Text())
    }
    /// e := echo.New()

    /// e.GET("/", cfilepath.Base(path))        //データの取得 アクセスに対して第2引き数（レスポンスの値）を返している
    /// :8080の後に
    //fmt.Println()

    // p22 より参照（bufio.Scanner）
    // 標準入力から読み込む
    //scanner := bufio.NewScanner(os.Stdin)
    // 1行ずつ読み込んで繰り返す
    //for scanner.Scan() {
        //1行分を出力する
        //      fmt.Fprintln(scanner.Text())
}
    //fmt.Println()


//
// 
//
//
//
// 引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンド
// 1.引数でファイルパスの一覧を貰う
/// コマンドライン（コンソールで入力する$ mycat -n hoge~）から文字列を受け取って 文字列を受け取る方法
//// 
/// 受け取った物を変数nに入れる 変数に入れる方法
/// ルート1 受け取った文字列をファイルパスに変換する 文字列を変換する方法

/// 引数から文字列（hogeとfuga）を読み取る（importでは無い）
//// 中でどういう処理をしているのか一つ

// 2.ファイルを順に標準出力に出力させる

// 3.行数
/// そもそも行数「$cat -n ファイル名」ってコマンドラインが行番号をつけて出力するので考慮しなくて良い。
/// "number"のn。



/* 使えそうなコマンド
・func flag.Args()[] 文字列
// フラグ以外のコマンドライン引数を返す

・func flag.Bool(○○ string, calue bool, usage string)* bool
// boolは真偽の型、
// (指定された名前、デフォルト値、使用方法の文字列)でブールフラグを定義する。
// ブールフラグの値を格納するブール変数のアドレスが戻り値として返ってきます。

・func flag.BoolVar（p * bool、name string、value bool、usage string）
// BoolVarは、(指定された名前、デフォルト値、使用方法の文字列)でブールフラグを定義する。
// 引数pは、フラグの値を格納するブール変数を指します。

・func Open（name string）（* File、error）
// Openは、名前付きファイルを読み取り用に開きます。
// 成功した場合、返されたファイルのメソッドを使用して読み取ることができます。
// 関連するファイル記述子のモードはO_RDONLYです。
// エラーがある場合、タイプは* PathErrorになります。

・type Scanner
//type Scanner struct {
//  // cフィルターされたフィールドまたはエクスポートされていないフィールドが含まれています
//}
// スキャナーは、改行で区切られたテキスト行のファイルなどのデータを読み取るための便利なインターフェイスを提供します。
// Scanメソッドへの連続した呼び出しは、トークン間のバイトをスキップして、ファイルの「トークン」をステップ実行します。
// トークンの指定は、タイプSplitFuncの分割関数によって定義されます。
// デフォルトの分割関数は、入力を行終端が取り除かれた行に分割します。
// このパッケージでは、ファイルを
// 行、バイト、UTF-8エンコードされたルーン、スペース区切りの
// 単語にスキャンするための分割関数が定義されています。
// クライアントは、代わりにカスタム分割関数を提供する場合があります。
// リーダーで順次スキャンを実行する必要があるプログラムは、代わりにbufio.Readerを使用する必要があります。

・bufio.scanner
//

・os.Stdout

・os.Stderr

・fmt.Fprintln
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
・


*/


/* 旧バージョン
import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// p23より参照
	// パスを結合する
	path := filepath.Join("dir", "main.go")
	fmt.Println(path)
	// ファイル名を取得
	fmt.Println(filepath.Base(path))

	// p22 より参照
	// 標準入力から読み込む
	scanner := bufio.NewScanner(os.Stdin)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		//1行分を出力する
		fmt.Println(scanner.Text())
	}
	fmt.Println()
}
*/
