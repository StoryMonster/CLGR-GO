
package search

import (
	"strings"
	"os"
	"bufio"
	"../common"
	"regexp"
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
	regs := make([]*regexp.Regexp, len(keywords))
	for i := 0; i < len(keywords); i++ {
		if ts.UseRegularMatch {
			regs[i], _ = regexp.Compile(keywords[i])
		} else if ts.IgnoreCase {
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
		lineCounter++
		for _, ch := range line { if isInvalidByte(ch) { return []common.MatchedLine{}, nil } }
		strLine := string(line)
		if len(keywords) == 0 {
			matchedLines = append(matchedLines, common.MatchedLine{lineCounter, strLine})
			continue
		}
		isFound := false
		if ts.UseRegularMatch {
			for _, reg := range regs {
				if reg.MatchString(strLine) {
					isFound = true
					break
				}
			}
		} else {
			if ts.IgnoreCase { strLine = strings.ToLower(strLine) }
			for _, keyword := range tempKeywords {
				// TODO 添加处理一行中有多个keyword的处理场景
				index := strings.Index(strLine, keyword)
				if index < 0 { continue }
				isFound = true
			    if ts.MatchWholeWord { isFound = common.IsWordWholeMatched(strLine, keyword, index) }
				break
			}
		}
		if isFound { matchedLines = append(matchedLines, common.MatchedLine{lineCounter, string(line)}) }
	}
	return matchedLines, nil
}