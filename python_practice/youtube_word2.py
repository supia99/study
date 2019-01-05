import re

n = open("new_sample2.txt", "w")#書き込みモードでファイル作成

with open("new_sample.txt", "r") as s:#読み込みモードで開く
    line = s.readline()
    i = 1
    while line:
        if i % 3 == 0 :
            n.write(line);
        if i == 5 :
            i = 1
        i = i + 1
        line = s.readline()

n.close()
