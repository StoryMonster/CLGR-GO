package search

import (
	"strings"
	"os"
	"bufio"
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
	if 
}

func (ts *TextSearcher)Search(keyword string) (matchedLines []MatchedLine, err error) {
	tempKeyword := keyword
	fd, err := os.Open(ts.DestFileName)
	if err != nil { return  }
	defer fd.Close()
	reader := bufio.NewReader(fd)
	for {
		line, _, err := reader.ReadLine()
		if err != nil { break }
		for 
	}

	if ts.UseRegularMatch {
		// 正则匹配
		return
	} else {
		if (ts.IgnoreCase) { tempKeyword = strings.ToLower(tempKeyword)}

	}
}