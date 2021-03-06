# 超A&G録音

## 概要
[超A&G+](https://www.agqr.jp/timetable/streaming.html)を予約録音する。  
以下の記事を参考にしている。  
https://k0bakatsu.hatenablog.com/entry/2020/11/12/013139

## 環境・使用条件
* Linux環境
* Cronが使用できる
* ffmpegがインストールされている  
  参照: https://ffmpeg.org/
* 録音したいタイミングにマシンが稼働している。

## 設計
* cronからスクリプト呼び出し
* スクリプト引数
    1. 番組名(半角スペースや半角記号が入らないようにする)
    1. 録音時間(録音する際は＋60sする)
    1. アップロード先ディレクトリ
