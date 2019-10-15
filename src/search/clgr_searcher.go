package search

import (
	"fmt"
	"../output"
	"os"
	"sync"
)

type ClgrSearcher struct {
	IgnoreCase bool                 // 若是文本搜索，只作用于文本匹配，不作用于文件查找
	MatchWholeWord bool             // 若是文本搜索，只作用于文本匹配，不作用于文件查找
	UseRegularMatch bool            // 若是文本搜索，同时作用于文本匹配和文件查找
	IgnoreFolderName bool           // 仅作用于文件名匹配
	DestDirs []string               // 搜索路径集合，路径必须完全正确
	DestFiles []string              // 文件名关键字集合
	TargetKeywords []string         // 仅用于文本搜索，作为文本搜索关键字集合
	op *output.Output
}

type TextSearchOptions struct {
	IgnoreCase bool         // 忽略大小写
	MatchWholeWord bool     // 全字匹配
	UseRegularMatch bool    // 正则匹配
	DestFileName string     // 待搜索文件名
}

func NewClgrSearcher(op *output.Output)(*ClgrSearcher, error) {
	return &ClgrSearcher{false, false, false, false, make([]string, 0), make([]string, 0), make([]string, 0), op}, nil
}

func (cs *ClgrSearcher)Search() {
	if len(cs.TargetKeywords) == 0 {
		cs.searchFiles()
	} else {
		cs.searchTexts()
	}
	cs.op.SearchConclusion()
}

func (cs *ClgrSearcher)searchFiles() {
    for _, dir := range cs.DestDirs {
		fs := NewFileSearcher(cs.IgnoreCase, cs.MatchWholeWord, cs.UseRegularMatch, cs.IgnoreFolderName, dir)
		filenames, err := fs.Search(cs.DestFiles)
		if err != nil {
			cs.op.ERROR(err.Error())
			continue
		}
		if len(filenames) == 0 {
			cs.op.WARN(fmt.Sprintf("No files found under %s", dir))
			continue
		}
		for _, filename := range filenames {
			cs.op.AddFileSearchResult(filename)
		}
	}
}

func (cs *ClgrSearcher)searchTexts() {
	var wg sync.WaitGroup
	for _, dir := range cs.DestDirs {
		fs := NewFileSearcher(true, false, cs.UseRegularMatch, true, dir)  // 文本搜索时，文件名匹配不考虑大小写以及全字匹配
		filenames, err := fs.Search(cs.DestFiles)
		if err != nil {
			cs.op.ERROR(err.Error())
			continue
		}
		if len(filenames) == 0 {
			cs.op.WARN(fmt.Sprintf("No files found under %s", dir))
			continue
		}
		for _, filename := range filenames {
			if IsTextFile(filename) {
				wg.Add(1)
				go cs.searchTextsInFile(&wg, filename)
			} else {
				cs.op.WARN(fmt.Sprintf("%s is not a text file", filename))
			}
		}
	}
	wg.Wait()
}

func (cs *ClgrSearcher)searchTextsInFile(wg *sync.WaitGroup, filename string) {
	defer wg.Done()
	keywords := cs.TargetKeywords
	ts := NewTextSearcher(cs.IgnoreCase, cs.MatchWholeWord, cs.UseRegularMatch, filename)
	matchedLines, err := ts.Search(keywords)
	if err != nil {
		cs.op.ERROR(err.Error())
		return
	}
	if len(matchedLines) == 0 {
		return
	}
	cs.op.AddTextSearchResult(ts.DestFileName, matchedLines)
	return
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

func (cs *ClgrSearcher)SetDestinationDirectors(dirs []string) {
	for _, dir := range dirs {
		_, err := os.Stat(dir)
        if err == nil || os.IsExist(err) {
			cs.DestDirs = append(cs.DestDirs, dir)
		} else {
			panic(fmt.Sprintf("%s is not exist!", dir))
		}
	}
}

func (cs *ClgrSearcher)SetTargetKeywords(keywords []string) {
	cs.TargetKeywords = keywords
}

func (cs *ClgrSearcher)SetDestinationFiles(filekeywords []string) {
	cs.DestFiles = filekeywords
}