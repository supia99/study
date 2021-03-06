# 超A&G録音

## 概要
[超A&G+](https://www.agqr.jp/timetable/streaming.html)を予約録音する。  
以下の記事を参考にしている。  
https://k0bakatsu.hatenablog.com/entry/2020/11/12/013139

## 環境・使用条件
* Linux環境
* Cronが使用できる
  * Cron稼働確認コマンド:`sudo service cron status`
* ffmpegがインストールされている  
  参照: https://ffmpeg.org/
* 録音したいタイミングにマシンが稼働している。

## 設計
* cronからスクリプト呼び出し
  * `/etc/cron.d`ディレクトリ以下にファイルを作ると良いかも
  例: ファイル名=radio-cron
  ~~~
SHELL=/bin/bash
PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin

# m h dom mon dow user  command
# 分 時 日　月 曜日 ユーザ コマンド
# 曜日: 0-6=日-土
30 16 * * 6 root /[スクリプトの配置ディレクトリ]/rec_AAndGPlus/recAAndGPlus.sh TrySailのTRYangle_harmony 1800 /[アップロード先ディレクトリ]/
  ~~~
  * cronの編集後は`sudo service cron restart`を実行する。
* スクリプト引数
    1. 番組名(半角スペースや半角記号が入らないようにする)
    1. 録音時間(録音する際は＋60sする)
    1. アップロード先ディレクトリ
