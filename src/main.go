package main

import (
	"./search"
)

func main() {
	destDirs := []string {"D:/projects", "D:/chunk_server_pc_mobile"}
	destFiles := []string {"main"}
	targetKeywords := []string {"int", "main"}
	ignoreCase := true
	matchWholeWord := true
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

}