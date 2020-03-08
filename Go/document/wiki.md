* P.30 goの実行可能ファイルはELFフォーマット(Linuxで一般的なフォーマット)ではなく、独自のフォーマットらしい
* P.38 フォルダに1つのパッケージ定義のみが原則
* P.39 "\_test.go"はパッケージをテストするためのファイルと認識される
* := の意味

* 組み込み関数: プログラミング言語であらかじめ使用できる関数

### 参照
* https://employment.en-japan.com/engineerhub/entry/2018/06/19/110000
* スターティングGo言語


---
* receiver https://go-tour-jp.appspot.com/methods/8
  * ポインタレシーバの使用理由
    1. メソッドがレシーバの指す変数を変更するため
    2. メソッドの呼び出しごとに変数のコピーを避けるため


* レシーバはメソッドを実行する構造体を決定するもの
  * しかし、Javaとは異なりメソッド、インスタンスがnullでもメソッドの実行は可能
    * どんな利点があるのか？
      * 利点: null(Goでは、nil)チェックをメソッド内に入れられる
        * Go言語でも、ただinterfaceの変数が宣言されただけでは,メソッド実行時にエラーを吐く
* interfaceはメソッドのシグニチャの集まり

* 空のインターフェース https://go-tour-jp.appspot.com/methods/14
  * 任意の型の値を保持できる
    * -> 未知の型を扱える

* スコープ: 頭文字が大文字であれば、他パッケージから参照可能
  * Javaでは、細かく指定できる

* １つの型で同じメソッド名を許さない

* 実装時に、スタックとヒープを気にする必要がない  
  * コンパイル時に自動で判別する
  * 変数への参照が残っていれば、ヒープに残る
    * https://qiita.com/rookx/items/a1e3d057a0ed71424094

---
## Go言語の概要
* Google社によって開発された言語
* 2009年に発表された
* 現在はオープンソース形式で開発されている
* シンプル
  * 予約語が少ない

* マルチプラットフォームで動作
  * Go言語では、クロスコンパイルによって別のOSでも単一ソースコードから実行ファイルを生成可能である
  * Javaでは、JVMによってOSの違いを埋めていた

---
## スライド
* qiita https://qiita.com/t-kusakabe/items/725e7438892bba395062
* github https://github.com/hakimel/reveal.js

---
## 情報
* slide https://www.slideshare.net/takuyaueda967/go-77689475
* goらしさ https://employment.en-japan.com/engineerhub/entry/2018/06/19/110000
* 読み直したいgo基礎 https://budougumi0617.github.io/2019/06/20/golangtokyo25-read-again-awesome-go-article/
* playground https://play.golang.org/
* Effective go http://go.shibu.jp/effective_go.html
* goつまづきポイント https://future-architect.github.io/articles/20190713/
