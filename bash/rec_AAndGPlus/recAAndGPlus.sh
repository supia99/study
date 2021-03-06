#!/bin/bash

#---
# 引数
#  1. 番組名
#  2. 録音時間
#  3. アップロード先ディレクトリ
#---

url=https://www.uniqueradio.jp/agplayer5/hls/mbr-0-cdn.m3u8

programName=$1
recTime=$2
uploadDir=$3

mkdir -p ~/tmp/radio

fileName=${programName}_`date "+%Y%m%d-%H%M%S".mp3`
ffmpeg -i ${url} -movflags faststart -t $((recTime+60)) ~/tmp/radio/${fileName}

cp ~/tmp/radio/${fileName} ${uploadDir}
