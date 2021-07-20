#!/bin/bash
#----
#処理
# ファイル名に入っている半角スペースを置き換えます。
#引数
# 1. 対象ディレクトリ
# 2. 置き換え後の文字列
#-----

if [ $# -ne 2 ] ; then
  echo "引数を与えてください。第１引数=対象ディレクトリ、第２引数＝置換後の文字"
fi

echo "対象ディレクトリ: "$1
echo "置換後文字列>>>"$2"<<<"
# read -p "Enterを押すと処理開始："

for file in ${1}*
do
  # echo "前:"$file
  # echo "後:"`echo $file | sed -e "s/ /$2/g"`
  rename=`echo $file | sed -e "s/ /$2/g"`
  if [ "$file" != $rename ] ; then
    mv "${file}" `echo $file | sed -e "s/ /$2/g"`
  fi
done

echo "ファイル名変更終了"
