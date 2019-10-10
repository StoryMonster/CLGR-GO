package search

import (
	"path/filepath"
	"os"
	"strings"
	"../common"
	"regexp"
)

type FileSearcher struct {
	IgnoreCase bool         // 忽略大小写
	MatchWholeWord bool     // 全字匹配
	UseRegularMatch bool    // 正则匹配
	IgnoreFolderName bool   // 不匹配文件夹名字
	DestDir string          // 搜索路径
}

func NewFileSearcher(ic bool, mww bool, urm bool, ifn bool, dd string) (*FileSearcher) {
	return &FileSearcher{ic, mww, urm, ifn, dd}
}

func (fs *FileSearcher)Search(keywords []string) (matchedFiles []string, err error) {
	tempKeywords := keywords
	regs := make([]*regexp.Regexp, len(keywords))
	for i := 0; i < len(keywords); i++ {
		if !fs.UseRegularMatch && fs.IgnoreCase {
			tempKeywords[i] = strings.ToLower(tempKeywords[i])
		} else {
			regs[i], _ = regexp.Compile(tempKeywords[i])
		}
	}

	filepath.Walk(fs.DestDir, func(path string, info os.FileInfo, err error) error {
		if fs.IgnoreFolderName && info.IsDir() { return err }
		filename := info.Name()
		if len(tempKeywords) == 0 {
			matchedFiles = append(matchedFiles, path)
			return err
		}
		isFound := false
		if !fs.UseRegularMatch {
			if fs.IgnoreCase {filename = strings.ToLower(filename) }
			for _, keyword := range tempKeywords {
				// TODO 添加处理一行中有多个keyword的处理场景
				index := strings.Index(filename, keyword)
				if index < 0 { continue }
				isFound = true
				if fs.MatchWholeWord { isFound = common.IsWordWholeMatched(filename, keyword, index) }
				break
			}
		} else {
			for i := 0; i < len(regs); i++ {
				isMatched := regs[i].MatchString(filename)
				if isMatched {
					isFound = true
					break
				}
			}
		}
		if isFound {
			matchedFiles = append(matchedFiles, path)
		}
		return err
	})
	err = nil
	return
}