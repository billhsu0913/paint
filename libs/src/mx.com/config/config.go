package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"mx.com/log"
)

var (
	NL  = []byte{'\n'}
	ANT = []byte{'#'}
)

func trimComments(data []byte) (data1 []byte) {

	conflines := bytes.Split(data, NL)
	for k, line := range conflines {
		conflines[k] = trimCommentsLine(line)
	}
	return bytes.Join(conflines, NL)
}

func trimCommentsLine(line []byte) []byte {

	var newLine []byte
	var i, quoteCount int
	lastIdx := len(line) - 1
	for i = 0; i <= lastIdx; i++ {
		if line[i] == '\\' {
			if i != lastIdx && (line[i+1] == '\\' || line[i+1] == '"') {
				newLine = append(newLine, line[i], line[i+1])
				i++
				continue
			}
		}
		if line[i] == '"' {
			quoteCount++
		}
		if line[i] == '#' {
			if quoteCount%2 == 0 {
				break
			}
		}
		newLine = append(newLine, line[i])
	}
	return newLine
}

func LoadEx(conf interface{}, confName string) (err error) {
	data, err := ioutil.ReadFile(confName)
	if err != nil {
		log.Error("Load conf file failed:", err)
		return
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		log.Error("Parse conf failed:", err)
	}
	return
}
