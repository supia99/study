#!/bin/bash

# ----
# 処理
# 　決まった期間ごとにファイルをTUFからLANDISKに移動させる
# 引数
# 　
# ---- 

# 録音データを読み込むNASのマウントディレクトリ
SOURCE_NAS=/mnt/radikoTufNas
# 録音データを読み込むディレクトリ
UPLOAD_SOURCE=${SOURCE_NAS}/Temp/
# 録音データをコピーする先のNASのマウントディレクトリ
DEST_NAS=/mnt/supirDirLandiskNas
# 録音データをコピーする先のディレクトリ
UPLOAD_DEST=${DEST_NAS}/radiko/
# 録音データでラジオ局がわからなかった場合のコピー先ディレクトリ
UPLOAD_OTHERS_DEST=${UPLOAD_DEST}その他/


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

## move
### ファイル名から半角スペース削除
bash $(cd $(dirname $0); pwd)/changeNameSpace.sh ${UPLOAD_SOURCE} 　

### アップロード先ファイル読み込み
#### 「ディレクトリ名,ファイル接頭辞」の形式
radioProgramList=`cat $(cd $(dirname $0); pwd)/list_radioProgram.csv`

#### ファイル名を配列で取得
ls -t --file-type ${UPLOAD_SOURCE} | grep -v "/" > ~/radioFileList
radioFileList=`cat ~/radioFileList`
rm ~/radioFileList

while read radioFile
do
  echo "start:"${radioFile}
  is_copied=false
  for radioProgram in ${radioProgramList}
  do
    radioStation=`echo ${radioProgram} | cut -f 1 -d ","`
    programName=`echo ${radioProgram} | cut -f 2 -d ","`
    # echo "|"${radioFile}"|"${programName}"|"
    if [[ ${radioFile} =~ ${programName} ]]; then

      uploadDest=${UPLOAD_DEST}${radioStation}"/"
      echo "sudo cp --preserve=timestamps --no-clobber \""${UPLOAD_SOURCE}${radioFile}"\" "${uploadDest}
      # sudo cp --preserve=timestamps --no-clobber ${UPLOAD_SOURCE}${radioFile} ${uploadDest}
      is_copied=true
      break
    fi
  done
  if ! "$is_copied"; then
    echo "存在しない："${radioFile}
    echo "sudo cp --preserve=timestamps --no-clobber \""${UPLOAD_SOURCE}${radioFile}"\" "${UPLOAD_OTHERS_DEST}
  fi
done << FILE
${radioFileList}
FILE


## コピーしなかったファイルを抽出する方法がわからず中止
### radioProgramNameごとにコピーする
# for radioProgram in ${radioProgramList}
# do
#   radioStation=`echo ${radioProgram} | cut -f 1 -d ","`
#   programName=`echo ${radioProgram} | cut -f 2 -d ","`
#   echo "|"${radioStation}"|"${programName}"|"
#   echo "cp --preserve=timestamps --no-clobber "${UPLOAD_SOURCE}${programName}"*" ${UPLOAD_DEST}
# done

# ### 上でコピーしなかったファイルをコピーする
# #### ファイル名を配列で取得
# ls -t --file-type ${UPLOAD_SOURCE} | grep -v "/" > ~/radioFileList
# radioFileList=`cat ~/radioFileList`
# rm ~/radioFileList
# while read radioFile
# do
#   echo "start:"${radioFile}
#   echo ${radioProgramList} | xargs -n 1 | grep -E ""
# done << FILE
# ${radioFileList}
# FILE



### findで15日以上前のファイルを削除する
echo `find ${UPLOAD_SOURCE} -maxdepth 1 -mtime +7`
