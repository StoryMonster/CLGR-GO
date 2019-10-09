package search

import (
	"path/filepath"
	"os"
	"strings"
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

func isSepareteCharactor(ch byte) bool {
	for _, item := range SPLITE_CHARACTORS {
		if ch == byte(item) {
			return true
		}
	}
	return false
}

func (fs *FileSearcher)Search(keywords []string) (matchedFiles []string, err error) {
	tempKeywords := keywords
	if !fs.UseRegularMatch && fs.IgnoreCase {
		for i := 0; i < len(keywords); i++ {
			tempKeywords[i] = strings.ToLower(tempKeywords[i])
		}
	}

	filepath.Walk(fs.DestDir, func(path string, info os.FileInfo, err error) error {
		isFound := false
		if !fs.UseRegularMatch {
			if fs.IgnoreFolderName && info.IsDir() { return err }
			filename := info.Name()
			if fs.IgnoreCase {filename = strings.ToLower(filename) }
			for _, keyword := range tempKeywords {
				index := strings.Index(filename, keyword)
				if index < 0 { continue }
				if fs.MatchWholeWord {
					leftIndex := index - 1
					rightIndex := index + len(keyword)
				    if leftIndex >= 0 && !strings.ContainsRune(SPLITE_CHARACTORS, rune(filename[leftIndex])) { continue }
					if rightIndex < len(filename) && !strings.ContainsRune(SPLITE_CHARACTORS, rune(filename[rightIndex])) { continue }
				}
				isFound = true
				break
			}
		} else {
			// TODO 正则匹配
		}
		if isFound { matchedFiles = append(matchedFiles, path) }
		return err
	})
	err = nil
	return
}