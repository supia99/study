import re

with open("Overwatch_Hero_Abilities_EVERYONE_Uses_WRONG.txt", "r") as s:#読み込みモードで開く
    text = s.read()

new_text = re.sub("<.*?>", "", text)

with open("new_sample.txt", "w") as n:#書き込みモードでファイル作成
    n.write(new_text)
