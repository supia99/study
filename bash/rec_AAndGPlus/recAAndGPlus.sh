#!/bin/bash

#---
# 引数
#  1. 番組名
#  2. 録音時間(秒)
#  3. アップロード先ディレクトリ
#---

url=https://www.uniqueradio.jp/agplayer5/hls/mbr-0-cdn.m3u8

# 録音データを読み込むNASのマウントディレクトリ
SOURCE_NAS=/mnt/radikoTufNas
# 録音データを読み込むディレクトリ
UPLOAD_SOURCE=${SOURCE_NAS}/Temp/
# 録音データをコピーする先のNASのマウントディレクトリ
DEST_NAS=/mnt/supirDirLandiskNas
# 録音データをコピーする先のディレクトリ
UPLOAD_DEST=${DEST_NAS}/radiko/
# 録音データでラジオ局がわからなかった場合のコピー先ディレクトリ
UPLOAD_OTHERS_AAndG=${UPLOAD_DEST}超A\&G/


## マウント用ディレクトリ作成.ある場合は、作成しない
if [ ! -d ${SOURCE_NAS} ]; then
  mkdir ${SOURCE_NAS}
fi
if [ ! -d ${DEST_NAS} ]; then
  mkdir ${DEST_NAS}
fi


## マウント
sudo mount -t drvfs '\\192.168.50.1\radiko' ${SOURCE_NAS}
sudo mount -t drvfs '\\landisk-c30161\supiaDir' ${DEST_NAS}



programName=$1
recTime=$2

mkdir -p ~/tmp/radio

fileName=${programName}_`date "+%Y%m%d-%H%M%S".mp3`
ffmpeg -i ${url} -movflags faststart -t $((recTime+60)) ~/tmp/radio/${fileName}

cp ~/tmp/radio/${fileName} ${UPLOAD_SOURCE}
cp ~/tmp/radio/${fileName} ${UPLOAD_OTHERS_AAndG}
