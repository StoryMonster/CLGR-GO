package main

import (
	"fmt"
	"./search"
)

func main() {
	fs := search.NewFileSearcher(false, false, false, true, ".")
	res := fs.Search("main")
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}