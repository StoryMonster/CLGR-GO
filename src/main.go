package main

import (
	"fmt"
	"./search"
)

func main() {
	/*
	fs := search.NewFileSearcher(false, false, false, true, ".")
	res := fs.Search("main")
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
	*/
	ts := search.NewTextSearcher(true, false, false, "./src/search/text_searcher.go")
	lines, err := ts.Search("Byte")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, line := range lines {
		fmt.Println(line.Line)
	}
}