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
		if err != nil { return err }
		if fs.IgnoreFolderName && info.IsDir() { return err }
		filename := info.Name()
		if len(tempKeywords) == 0 {
			matchedFiles = append(matchedFiles, path)
			return err
		}
		if !fs.UseRegularMatch {
			if fs.IgnoreCase {filename = strings.ToLower(filename) }
			for _, keyword := range tempKeywords {
				tempFilename := filename
				for ; len(tempFilename) > 0;{
					index := strings.Index(tempFilename, keyword)
					if index < 0 { break }
					isFound := true
					if fs.MatchWholeWord { isFound = common.IsWordWholeMatched(tempFilename, keyword, index) }
					if isFound {
						matchedFiles = append(matchedFiles, path)
						return err
					}
					tempFilename = tempFilename[index+len(keyword):]
				}
			}
		} else {
			for i := 0; i < len(regs); i++ {
				if regs[i].MatchString(filename) {
					matchedFiles = append(matchedFiles, path)
					return err
				}
			}
		}
		return err
	})
	err = nil
	return
}