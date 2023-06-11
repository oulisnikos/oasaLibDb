package oasaSyncDecode

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

const Space = 32
const DoubleQuotes = 34
const ComaCharcter = 44
const OpenParenthesis = 40
const CloseParenthesis = 41

func ReadRecCharByCar(recordStr string) []string {
	var result []string
	r := bufio.NewReader(strings.NewReader(recordStr))

	var strStartStop bool = false
	var addCharacter = false
	var prop = ""
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			charecter := string(c)

			if charecter == string(DoubleQuotes) {
				strStartStop = !strStartStop
			}

			if strStartStop && charecter != string(DoubleQuotes) {
				addCharacter = true
			} else if !strStartStop && charecter != string(DoubleQuotes) {
				if charecter == string(Space) {

				} else if charecter == string(OpenParenthesis) {
					//continue
				} else if charecter == string(CloseParenthesis) {
					result = append(result, prop)
					//continue
				} else if charecter == string(ComaCharcter) {
					result = append(result, prop)
					prop = ""
					//continue
				} else {
					addCharacter = true
				}
			}

			if addCharacter {
				addCharacter = false
				prop = prop + charecter
			}

		}
	}
	return result
}

func ReadTextCharByChar(inputStr string, mapToDbRec func(dataArray []string) map[string]interface{}, file *os.File) []map[string]interface{} {
	var resultData []map[string]interface{}
	r := bufio.NewReader(strings.NewReader(inputStr))
	//var i = 0
	//f, err := os.Create("/oasa-telematics/data.txt")

	//if err != nil {
	//	fmt.Printf("Error in Open File [%s]", err.Error())
	//}

	//defer f.Close()
	var recStartStop = false
	var quotedStr = false
	var recStr = ""
	var i = 0
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if i == 641 {
				var ii = i
				i = i + 1
				i = ii
			}
			charecter := string(c)
			if charecter == string(DoubleQuotes) {
				quotedStr = !quotedStr
			} else if charecter == string(OpenParenthesis) && !quotedStr {
				recStartStop = true
			} else if charecter == string(CloseParenthesis) && !quotedStr {
				recStartStop = false
				recStr = recStr + charecter
				file.WriteString(recStr + "\n")
				propArray := ReadRecCharByCar(recStr)
				if mapToDbRec != nil {
					resultData = append(resultData, mapToDbRec(propArray))
				}
				//fmt.Print("These props of rec ", propArray, " \n")
				recStr = ""
				i++
				continue
			}
			if recStartStop {
				recStr = recStr + charecter
			}
		}
	}
	return resultData
}
