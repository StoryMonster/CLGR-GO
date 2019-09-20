package search

import (
	"path/filepath"
	"os"
	"strings"
	"fmt"
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

func (fs *FileSearcher)Search(keyword string) (matchedFiles []string, err error) {
	tempKeyword := keyword
	if !fs.UseRegularMatch && fs.IgnoreCase { tempKeyword = strings.ToLower(tempKeyword) }

	filepath.Walk(fs.DestDir, func(path string, info os.FileInfo, err error) error {
		if !fs.UseRegularMatch {
			if fs.IgnoreFolderName && info.IsDir() { return err }
			filename := info.Name()
			if fs.IgnoreCase {filename = strings.ToLower(filename) }
			index := strings.Index(filename, tempKeyword)
			if index < 0 { return err }
		    if fs.MatchWholeWord {
				spliteChars := " -+_\\/><*&^%$#!@~`?[]{}()=|:;,."
				leftIndex := index - 1
				rightIndex := index + 1
				if leftIndex > 0 {
					if strings.IndexByte(spliteChars, filename[leftIndex]) < 0 { return err }
				}
				if rightIndex < len(filename) {
                    if strings.IndexByte(spliteChars, filename[rightIndex]) < 0 { return err }
				}
				matchedFiles = append(matchedFiles, path)
				fmt.Println(path)  // 用于快速输出结果到屏幕上
			}
		} else {
			// TODO 正则匹配
		}
		matchedFiles = append(matchedFiles, path)
		return err
	})
	err = nil
	return
}