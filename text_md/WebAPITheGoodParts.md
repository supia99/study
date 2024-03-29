## 2.3 httpメソッド
URIとメソッドの関係は、操作するものと操作方法の関係です。
URIが操作するもの(リソース)、HTTPメソッドが操作方法を示す。
　　(HTML4.0で、FormタグがGETとPOSTしか使うことができないため、通常はGETとPOSTだけを利用する場合が多い)

使い分けは以下。
* GET ... リソースの取得
* POST ... リソースの新規登録
* PUT ... 既存リソースの更新
* PATCH ... 既存リソースの一部変更
* DELETE ... 既存リソースの一部変更
(* HEAD ... リソースのメタ情報の取得)

## 2.4 APIのエンドポイント設計


## 2.5 検索とクエリパラメータの設計
クエリパラメータの使用例: http://api.example.com/user?sex=male

クエリパラメータは絞り込みに使用する
* 検索条件
* ページネーション

### 2.5.1 - 2.5.2 相対的な位置によるページネーション
相対的=データの前から何番目から何番目で取得する方法。
データ内の相対的な位置によりページネーションする方法には、2つのパターンがある
* per_page & page ... 1ページ当たりの個数と何ページ目を取得するかを指定する
* limit & offset ... 取得上限数と先頭から1つ目から何個を飛ばしてから取得するかを指定する
per_page=50&page=3 と limit=50&offset=100　は同じ意味。
limit&offsetの方がどこから取得するかを選べる点で自由度が高い分想定外のアクセスが増えるのでキャッシュ効率が下がる可能性がある。
API内でどちらを使うかを統一していればよい。

相対的な位置による方法の問題点は2つ。
* 前から数える処理があることで、遅くなる
* 更新頻度が多い場合に、データ不整合が起きやすい

### 2.5.3 絶対的な位置によるページネーション
絶対的=データのプロパティ(ID等)により一部を取得する方法。
相対的な位置によるページネーションで発生していた問題は解決する。
しかし、決まった個数を取得するのは難しくなりやすい？

### 2.5.4 絞り込みのためのパラメータ
いくつのかのクエリパラメータを使って絞り込んで(=検索)
1つのクエリパラメータのみで絞り込む、「q」というパラメータ名を使う場合がある。
「q」をクエリパラメータとして使う場合、一般的に部分一致で検索する意味を指すことが多いし、全カラムの検索を意味することもある。(色々な使い方がされている)

### 2.5.5 URIに単語「search」を含めるか
全件取得ではなく、検索を意識させたい場合には「search」を入れるのはアリ。

### 2.5.6 クエリパラメータとパスの使い分け
パスに含める基準は2つ
* 一意なリソースを表すのに必要な情報である
* 省略不可である


## 2.6 ログインとOAuth 2.0
認証での一般的な標準仕様はOAuthです。
他サービスの認証を使って、認証させる方法。





