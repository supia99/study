#!/bin/bash

#---
# 引数
#  1. 番組名
#  2. 録音時間(秒)
#---

url=https://www.uniqueradio.jp/agplayer5/hls/mbr-0-cdn.m3u8

# TUF NAS のアップロード先
UPLOAD_TUF=/mnt/radikoTufNas/Temp/

# 録音データでラジオ局がわからなかった場合のコピー先ディレクトリ
UPLOAD_LANDISK=/mnt/supirDirLandiskNas/radiko/超A\&G/


programName=$1
recTime=$2

mkdir -p ~/tmp/radio

fileName=${programName}_`date "+%Y%m%d-%H%M%S".mp3`
ffmpeg -i ${url} -movflags faststart -t $((recTime+60)) ~/tmp/radio/${fileName}

cp ~/tmp/radio/${fileName} ${UPLOAD_TUF}
cp ~/tmp/radio/${fileName} ${UPLOAD_LANDISK}
