package search

import (
	"strings"
	"os"
	"bufio"
	"errors"
)

type MatchedLine struct {
	Num int     // 行号
	Line string // 行内容
}

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

func (ts *TextSearcher)Search(keyword string) (matchedLines []MatchedLine, err error) {
	tempKeyword := keyword
	if ts.IgnoreCase && !ts.UseRegularMatch { tempKeyword = strings.ToLower(tempKeyword) }

	fd, err := os.Open(ts.DestFileName)
	if err != nil { return  }
	defer fd.Close()
	reader := bufio.NewReader(fd)
	lineCounter := 0
	for {
		line, _, err := reader.ReadLine()
		if err != nil { break }
		lineCounter++
		for _, ch := range line {
			if isInvalidByte(ch) {
				return matchedLines, errors.New("Read text from a binary file")
			}
		}
		if ts.UseRegularMatch {
			// 正则匹配
		    return matchedLines, nil
		} else {
			strLine := string(line)
			if ts.IgnoreCase { strLine = strings.ToLower(strLine) }
			index := strings.Index(strLine, tempKeyword)
			if index < 0 { continue }
			if ts.MatchWholeWord {
				spliteChars := " -+_\\/><*&^%$#!@~`?[]{}()=|:;,."
				leftIndex := index - 1
				rightIndex := index + 1
				if leftIndex > 0 {
					if strings.IndexByte(spliteChars, strLine[leftIndex]) < 0 { continue }
				}
				if rightIndex < len(strLine) {
                    if strings.IndexByte(spliteChars, strLine[rightIndex]) < 0 { continue }
				}
			}
			matchedLines = append(matchedLines, MatchedLine{lineCounter, string(line)})
		}
	}
	if len(matchedLines) == 0 {
		return matchedLines, errors.New("No text found!")
	}
	return matchedLines, nil
}