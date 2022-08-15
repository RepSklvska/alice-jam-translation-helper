package main

import (
	"bufio"
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

func ReadDialogFileToJson(path string) error {
	var lines, err = ReadLines(path)
	if err != nil {
		return err
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
		//fmt.Println(line)
		if strings.Contains(lines[i], "CALLFUNC NAME") {
			lineNumberOfName := i - 1
			name := func() string {
				// “CALLFUNC NAME”的上一行的引号里就是角色名称
				lineContainsName := lines[i-1]
				index1 := strings.Index(lineContainsName, "\"") + 1
				index2 := strings.LastIndex(lineContainsName, "\"")
				return lines[i][index1:index2]
			}()

			var (
				lineNumberOfMessageContent int
				messageContent             []string
			)
			for ii := i; ; i++ {
				if strings.Contains(lines[ii], " A") {
					// 如果结尾是“ A”，则得知这是最后一行，往前找有没有结尾是“ R”的

				}
			}
			//dialog.Messages = append(dialog.Messages)
		}
	}

	return nil
}

func main() {
	os.Chdir(BasePath)
	os.Chdir(ScenarioTextPath)
	ReadDialogFileToJson(TestFile)
}
