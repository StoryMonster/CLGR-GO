package search

var SPLITE_CHARACTORS = " -+_\\/><*&^%$#!@~`?[]{}()=|:;,."

var BINARY_FILES = []string {
	// libs
	".dll", ".so", ".a",
	// binary files
	".exe", ".out", ".bin", ".apk", ".msi",
	// microsoft office
	".doc", ".pptx", ".ppt", ".xlsm", ".pdf", ".xlsx", ".docx",
	// video
	".mp4", ".avi", ".rmvb",
	// audio
	".mp3", ".wmv",
	// compress
	".tar", ".zip", ".rar",
	// others
	".o", ".pyc"}