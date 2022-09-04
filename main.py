BasePath = "C:\\Users\\ekito\\Desktop\\LO-TMod-195\\"
ScenarioTextPath = "HT僔僫儕僆\\僔僫儕僆杮懱\\"
TestFile = "杮曇1\\僔僫儕僆HTLO_S001_A.jam"
SampleFile = "samples\\195\\僔僫儕僆HTLO_S001_A.jam"

import json


def ExtractString(str, char):
	str = str.strip()
	try:
		lIndex = str.index(char) + 1
		rIndex = str.rindex(char)
	except ValueError:
		return ""
	else:
		return str[lIndex:rIndex]


def ReadJamDialogToInfo(filename):
	file = open(filename, "r", encoding="utf-8")
	lines = file.readlines()
	dialog = []
	for i in range(0, len(lines)):
		if lines[i].__contains__("CALLFUNC NAME"):
			# 角色的名字
			characterName = ExtractString(lines[i - 1], "\"")
			# print("---", i)
			# print(characterName)
			sentences = []
			for ii in range(i, len(lines)):
				if lines[ii].strip()[-2:] == " A":
					# 如果结尾是“ A”，则得知这是最后一行，往前找有没有结尾是“ R”的
					lastLine = ExtractString(lines[ii].strip(), lines[ii].strip()[-3])
					# print(ii, lastLine)
					sentences.append(lastLine)
					for iii in range(ii, i, -1):
						if lines[iii].strip()[-2:] == " R":
							thisLine = ExtractString(lines[iii].strip(), lines[iii].strip()[-3])
							# print(iii, thisLine)
							sentences.insert(0, thisLine)
					break
			# print(sentences)
			dialog.append({characterName: sentences})
	# print(dialog)
	indexLastSentence = 0
	for i in range(len(lines) - 1, 0, -1):
		if len(lines[i].strip()) >= 2 and lines[i].strip()[-2:] == " A":
			indexLastSentence = i + 1
			break
	file.close()
	return {"dialog": dialog, "last": indexLastSentence}


sampleDialogInfo = ReadJamDialogToInfo(SampleFile)
# print(sampleDialogInfo)
jsonInfo = json.dumps(sampleDialogInfo, indent="\t", sort_keys=True, ensure_ascii=False)
print(jsonInfo)


# 读取单个文件的已经完成,剩下的是用这个文件写入的过程

def ReadJsonRewriteJam():
	a = 114514
