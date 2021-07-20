#!/bin/bash

#---
# sudo mount -t drvfs '\\192.168.50.1\radiko' /mnt/nasRadiko
# 参照: https://superuser.com/questions/1491857/windows-share-folder-mount-error-cifs-filesystem-not-supported-by-the-system
# 認証情報はWindows資格情報を使用する
# WSL2上でのみ動く。Linuxで動かす場合は、nasの認証について改修が必要。
#
# 引数
# 第１引数: allの場合、アップロード元のソースを元に全てのファイルのコピーを試みる
#          すでにアップロード先にファイルが存在する場合は、コピーしない。
#---

# 引数チェック
if [ "$1" == "all" ]; then
  IS_ALL=true
else
  IS_ALL=false
fi


# アップロード先・アップロード元定数定義
## アップロード先・アップロード元定数定義ファイル
DEFINE_FILE=~/radiko_copy.csv

if [ ! -f ${DEFINE_FILE} ]; then
  echo "定義ファイルが存在しません:"${DEFINE_FILE}
  echo "例: UPLOAD_SOURCE,/mnt/f/amusement/radio_rec/,"
  echo "UPLOAD_DEST,/mnt/nasRadiko/Temp/,"
  exit 1
fi

# アップロード先・アップロード元定数定義ファイル読み込み
for uploadInfo in `cat ${DEFINE_FILE}`
do
  if [ `echo ${uploadInfo} | cut -f 1 -d ","` == "UPLOAD_SOURCE" ]; then
    UPLOAD_SOURCE=`echo ${uploadInfo} | cut -f 2 -d ","`
    echo "UPLOAD_SOURCE="${UPLOAD_SOURCE}
  elif [ `echo ${uploadInfo} | cut -f 1 -d ","` == "UPLOAD_DEST" ]; then
    UPLOAD_DEST=`echo ${uploadInfo} | cut -f 2 -d ","`
    echo "UPLOAD_DEST="${UPLOAD_DEST}
  fi
done

## 変数が定義されているかの確認
if [ -z ${UPLOAD_SOURCE} -o -z ${UPLOAD_DEST} ]; then
  echo "定義ファイルがにUPLOAD_SOURCEかUPLOAD_DESTが定義されていません。:"${DEFINE_FILE}
  exit 1
fi
echo "upload:" ${UPLOAD_SOURCE} ${UPLOAD_DEST}

# UPLOAD_SOURCE="/mnt/f/amusement/radio_rec/"
# UPLOAD_DEST="/mnt/nasRadiko/Temp/"


# NASマウント
sudo mount -t drvfs '\\192.168.50.1\radiko' /mnt/radikoTufNas

# ファイル名の半角スペースを置換
bash $(cd $(dirname $0); pwd)/changeNameSpace.sh ${UPLOAD_SOURCE}


# アップロード先ファイル読み込み
## 「ディレクトリ名,ファイル接頭辞」の形式
radioProgramList=`cat $(cd $(dirname $0); pwd)/list_radioProgram.csv`
# for radioProgram in ${radioProgramList}
# do
#   radioStation=`echo ${radioProgram} | cut -f 1 -d ","`
#   programName=`echo ${radioProgram} | cut -f 2 -d ","`
#   # echo "1:"${radioStation}" 2:"${programName}
# done


# 音声ファイル名を取得
find ${UPLOAD_SOURCE} -maxdepth 1 -mtime -14 > ~/radioFileList
radioFileList=`tail -n +2 ~/radioFileList`

# https://server.etutsplus.com/sh-while-read-line-4pattern/
# while read radioFile
# do
#   echo "echo:"$radioFile
# done << FILE
# $radioFileList
# FILE

while read radioFilePath
do
  radioFileName=`echo ${radioFilePath} | sed -e "s;${UPLOAD_SOURCE};;"`
  if [[ `find ${UPLOAD_DEST} -name ${radioFileName}` ]]; then 
    sizeDest=`ls ${UPLOAD_DEST}${radioFileName} -l | cut -f 5 -d " "`
    sizeSource=`ls ${radioFilePath} -l | cut -f 5 -d " "`
    echo ${sizeSource} ${sizeDest}
    if [[ ${sizeDest} -ge ${sizeSource} ]]; then
      echo ${radioFileName}は既に存在する。かつ、ファイルサイズがそれ以下ですので、コピーしませんでした。
      continue
    fi  
  fi
  echo cp -f --preserve=timestamps ${radioFilePath} ${UPLOAD_DEST}
  ## 調査：タイムスタンプが保持されない原因
  cp -p ${radioFilePath} ${UPLOAD_DEST}
done << FILE
${radioFileList}
FILE
