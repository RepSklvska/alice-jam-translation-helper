package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//
var (
	BasePath         = "C:\\Users\\ekito\\Desktop\\LO-TMod-195\\"
	ScenarioTextPath = "HT僔僫儕僆\\僔僫儕僆杮懱\\"

	TestFile = "杮曇1\\僔僫儕僆HTLO_S001_A.jam"
)

func ReadLines(path string) ([]string, error) {
	var lines []string

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return lines, err
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

func GetTextFromLine(line string) (text string) {
	line = strings.TrimSpace(line)
	text = line[2 : len(line)-3]
	return text
}

func ReadDialogFileToJson(path string) (string, error) {
	var lines, err = ReadLines(path)
	if err != nil {
		return "{}", err
	}

	// 输出JSON的结构
	type Content struct {
		Text []string `json:"text"`
		Line int      `json:"line"`
	}
	type Message struct {
		Name    string  `json:"name"`
		Line    int     `json:"line"`
		Message Content `json:"message"`
	}
	type Dialog struct {
		Messages []Message `json:"messages"`
	}

	var dialog *Dialog = new(Dialog)
	for i := 0; i < len(lines); i++ {
		//fmt.Println(lines[i])
		if strings.Contains(lines[i], "CALLFUNC NAME") {
			//fmt.Println(lines[i-1])
			lineNumberOfName := i - 1
			name := func() string {
				// “CALLFUNC NAME”的上一行的引号里就是角色名称
				lineContainsName := lines[i-1]
				index1 := strings.Index(lineContainsName, "\"") + 1
				index2 := strings.LastIndex(lineContainsName, "\"")
				return lineContainsName[index1:index2]
			}()
			fmt.Println(name)
			var (
				lineNumberOfMessageContent int
				messageContent             []string = []string{}
			)
			for ii := i; ; ii++ {
				if strings.LastIndex(lines[ii], " A") == len(lines[ii])-2 {
					// 如果结尾是“ A”，则得知这是最后一行，往前找有没有结尾是“ R”的
					messageContent = append(messageContent, GetTextFromLine(lines[ii]))
					//fmt.Println(lines[ii])
					fmt.Println(messageContent)
					for iii := ii - 1; ; iii-- {
						if strings.LastIndex(lines[iii], " R") == len(lines[iii])-2 {
							//fmt.Println(lines[iii])
							messageContent = append([]string{GetTextFromLine(lines[iii])}, messageContent...)
							fmt.Println(messageContent)
						} else {
							lineNumberOfMessageContent = iii + 1
							break
						}
					}
					break
				}
			}
			var message *Message = new(Message)
			message.Line = lineNumberOfName
			message.Name = name
			message.Message.Line = lineNumberOfMessageContent
			message.Message.Text = messageContent
			dialog.Messages = append(dialog.Messages)
		}
	}
	fmt.Println(dialog)
	//jsonBytes, err := json.Marshal(dialog)

	return "", nil
}

func main() {
	//os.Chdir(BasePath)
	//os.Chdir(ScenarioTextPath)
	ReadDialogFileToJson("./samples/195/僔僫儕僆HTLO_S001_A.jam")
}
