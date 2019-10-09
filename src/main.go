package main

import (
	"./search"
	"fmt"
)

func main() {
	/*
	destDirs := []string {"D:/projects", "D:/chunk_server_pc_mobile"}
	destFiles := []string {"main"}
	targetKeywords := []string {"int", "main"}
	ignoreCase := false
	matchWholeWord := false
	useRegularMatch := false
	ignoreFolderName := true

	cs, _ := search.NewClgrSearcher()
	cs.SetIgnoreCase(ignoreCase)
	cs.SetMatchWholeWord(matchWholeWord)
	cs.SetUseRegularMatch(useRegularMatch)
	cs.SetIgnoreFolderName(ignoreFolderName)
	cs.SetDestinationDirectors(destDirs)
	cs.SetTargetKeywords(targetKeywords)
	cs.SetDestinationFiles(destFiles)

	cs.Search()
	*/
	destDirs := []string {"D:/projects"}
	destFiles := []string {"main"}
	//targetKeywords := []string {"int", "main"}
	ignoreCase := false
	matchWholeWord := true
	useRegularMatch := false
	ignoreFolderName := true

	fs := search.NewFileSearcher(ignoreCase, matchWholeWord, useRegularMatch, ignoreFolderName, destDirs[0])
	matchedFiles, _ := fs.Search(destFiles)
	for _, filename := range matchedFiles {
		fmt.Println(filename)
	}
}