
package search

import (
	"strings"
	"os"
	"bufio"
	"../common"
)

type TextSearcher struct {
	IgnoreCase bool         // 忽略大小写
	MatchWholeWord bool     // 全字匹配
	UseRegularMatch bool    // 正则匹配
	DestFileName string     // 待搜索文件名
}

func NewTextSearcher(ic bool, mww bool, urm bool, dfn string) (*TextSearcher) {
	return &TextSearcher{ic, mww, urm, dfn}
}

func isInvalidByte(bt byte) bool {
	return bt >= 0x00 && bt <= 0x08
}

func IsTextFile(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil { return false }
	if info.IsDir() { return false }

	// 仅从文件名的角度进行检查
	index := strings.LastIndex(filename, ".")
	if index < 0 { return true }
	suffix := filename[index:]
	for _, item := range common.BINARY_FILES {
		if strings.Compare(item, suffix) == 0 {
			return false
		}
	}
	return true
}

func (ts *TextSearcher)Search(keywords []string) (matchedLines []common.MatchedLine, err error) {
	tempKeywords := keywords
	if ts.IgnoreCase && !ts.UseRegularMatch {
		for i := 0; i < len(keywords); i++ {
			tempKeywords[i] = strings.ToLower(tempKeywords[i])
		}
	}

	fd, err := os.Open(ts.DestFileName)
	if err != nil { return  }
	defer fd.Close()
	reader := bufio.NewReader(fd)
	lineCounter := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil { break }
		isFound := false
		lineCounter++
		for _, ch := range line {
			if isInvalidByte(ch) {
				return []common.MatchedLine{}, nil
			}
		}
		if ts.UseRegularMatch {
			// 正则匹配
		} else {
			strLine := string(line)
			if ts.IgnoreCase { strLine = strings.ToLower(strLine) }
			if len(tempKeywords) == 0 { isFound = true }
			for _, keyword := range tempKeywords {
				index := strings.Index(strLine, keyword)
			    if index < 0 { continue }
			    if ts.MatchWholeWord {
					leftIndex := index - 1
					rightIndex := index + len(keyword)
				    if leftIndex >= 0 && !strings.ContainsRune(common.SPLITE_CHARACTORS, rune(strLine[leftIndex])) { continue }
					if rightIndex < len(strLine) && !strings.ContainsRune(common.SPLITE_CHARACTORS, rune(strLine[rightIndex])) { continue }
				}
				isFound = true
				break
			}
			if isFound { matchedLines = append(matchedLines, common.MatchedLine{lineCounter, string(line)}) }
		}
	}
	return matchedLines, nil
}