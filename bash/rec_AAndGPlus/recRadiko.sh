#!/bin/bash

#---
# 引数
#　1. ラジオ局
#  2. 番組名
#  3. 録音時間(分)
#---

# TUF NAS のアップロード先
UPLOAD_TUF=/mnt/radikoTufNas/Temp/

# 録音データでラジオ局がわからなかった場合のコピー先ディレクトリ
UPLOAD_LANDISK=/mnt/supirDirLandiskNas/radiko/

radioStation=$1
programName=$2
recTime=$3

mkdir -p ~/tmp/radio

fileName=${programName}_`date "+%Y%m%d-%H%M%S"`
/home/kawajiri/github/study/bash/rec_AAndGPlus/radish/radi.sh -t radiko -s ${radioStation} -d $((recTime+1)) -o ~/tmp/radio/${fileName}

cp ~/tmp/radio/${fileName}* ${UPLOAD_TUF}
sudo cp ~/tmp/radio/${fileName}* ${UPLOAD_LANDISK}${radioStation}/
