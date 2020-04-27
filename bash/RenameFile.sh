# !/bin/bash
#49 150306

# テスト環境を戻す
rm ichigo*
cp tmp/* ./
# 実施
for file in ichigo* ;
do
  renameFile=$(echo $file | sed -e "s/[0-9]\+_/_/");
  echo rename $file "->" $renameFile;
  mv $file $renameFile;
done

fileNum=1
for file in ichigo* ;
do
  fileNumFormatted=`printf %03d_ $fileNum`
  echo $fileNumFormatted $file;
  renameFile=$(echo $file | sed -e "s/_/$fileNumFormatted/")
  echo "renamefilename:" $renameFile
  mv $file $renameFile
  fileNum=$((fileNum+1))
done
