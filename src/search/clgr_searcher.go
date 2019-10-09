package search

import (
	"fmt"
)

type ClgrSearcher struct {
	IgnoreCase bool                 // 若是文本搜索，只作用于文本匹配，不作用于文件查找
	MatchWholeWord bool             // 若是文本搜索，只作用于文本匹配，不作用于文件查找
	UseRegularMatch bool            // 若是文本搜索，同时作用于文本匹配和文件查找
	IgnoreFolderName bool           // 仅作用于文件名匹配
	DestDirs []string               // 搜索路径集合，路径必须完全正确
	DestFiles []string              // 文件名关键字集合
	TargetKeywords []string         // 仅用于文本搜索，作为文本搜索关键字集合
}

type TextSearchOptions struct {
	IgnoreCase bool         // 忽略大小写
	MatchWholeWord bool     // 全字匹配
	UseRegularMatch bool    // 正则匹配
	DestFileName string     // 待搜索文件名
}

func NewClgrSearcher()(*ClgrSearcher, error) {
	return &ClgrSearcher{false, false, false, false, make([]string, 0), make([]string, 0), make([]string, 0)}, nil
}

func (cs *ClgrSearcher)Search() {
	isTextSearch := len(cs.TargetKeywords) > 0
	if !isTextSearch {
		cs.searchFiles()
	} else {
		cs.searchTexts()
	}
}

func (cs *ClgrSearcher)searchFiles() {
    for _, dir := range cs.DestDirs {
		fs := NewFileSearcher(cs.IgnoreCase, cs.MatchWholeWord, cs.UseRegularMatch, cs.IgnoreFolderName, dir)
		filenames, err := fs.Search(cs.DestFiles)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if len(filenames) == 0 {
			fmt.Println(fmt.Sprintf("[WARN] No files found under %s", dir))
			continue
		}
		for _, filename := range filenames {
			fmt.Println(filename)
		}
	}
}

func (cs *ClgrSearcher)searchTexts() {
	for _, dir := range cs.DestDirs {
		fs := NewFileSearcher(true, false, cs.UseRegularMatch, cs.IgnoreFolderName, dir)  // 文本搜索时，文件名匹配不考虑大小写以及全字匹配
		filenames, err := fs.Search(cs.DestFiles)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if len(filenames) == 0 {
			fmt.Println(fmt.Sprintf("[WARN] No files found under %s", dir))
			continue
		}
		for _, filename := range filenames {
			if IsTextFile(filename) {
				ts := NewTextSearcher(cs.IgnoreCase, cs.MatchWholeWord, cs.UseRegularMatch, filename)
			    go searchTextsInFile(ts, cs.TargetKeywords)
			} else {
				fmt.Println(fmt.Sprintf("[WARN] %s is a binary file", filename))
			}
		}
	}
}

func searchTextsInFile(ts *TextSearcher, keywords []string) {
	matchedLines, err := ts.Search(keywords)
	if err != nil {
		fmt.Println(fmt.Sprintf("[WARN] %s", err.Error()))
		return
	}
	if len(matchedLines) == 0 {
		return
	}
	res := fmt.Sprintf(">>>%s:\n", ts.DestFileName)
	for _, line := range matchedLines {
		res += fmt.Sprintf("%d: %s\n", line.Num, line.Line)
	}
	fmt.Println(res)
}

func (cs *ClgrSearcher)SetIgnoreCase(val bool) {
	cs.IgnoreCase = val
}

func (cs *ClgrSearcher)SetMatchWholeWord(val bool) {
	cs.MatchWholeWord = val
}

func (cs *ClgrSearcher)SetUseRegularMatch(val bool) {
    cs.UseRegularMatch = val
}

func (cs *ClgrSearcher)SetIgnoreFolderName(val bool) {
    cs.IgnoreFolderName = val
}

func (cs *ClgrSearcher)SetDestinationDirectors(paths []string) {
	cs.DestDirs = paths
}

func (cs *ClgrSearcher)SetTargetKeywords(paths []string) {
	cs.TargetKeywords = paths
}

func (cs *ClgrSearcher)SetDestinationFiles(paths []string) {
	cs.DestFiles = paths
}