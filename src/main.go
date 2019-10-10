package main

import (
	"./search"
	"./output"
	"./args"
	"os"
)

func readCmdParameters() *args.Args {
	arg := args.New("CLGR", "v1.0")
	arg.AddParameter("text", "t", []string {}, "Text search keyword, multiple text can be used.")
	arg.AddParameter("file", "f", []string {}, "File search keyword, multiple file can be used.")
	arg.AddParameter("dir", "d", []string {"."}, "Specify the search directory, multiple dir can be used.")
	arg.AddParameter("ignorecase", "ic", []string {"false"}, "Ignore the alphabet case.")
	arg.AddParameter("matchwholeword", "mwc", []string {"false"}, "Match the whole word in a line or a filename when search.")
	arg.AddParameter("ignorefoldername", "ifn", []string {"false"}, "Ignore the name of folder when do search for files.")
	arg.AddParameter("regular", "r", []string {"false"}, "Use regular expression to do match")
	arg.Parse()
	return arg
}

func noExpectParamExist(arg *args.Args) bool {
    for _, param := range arg.Parameters {
		if param.IsExist { return false }
	}
	return true
}

func main() {
	arg := readCmdParameters()
	if noExpectParamExist(arg) { return }
	op := output.New(os.Stdout, os.Stdout, output.Error)
	cs, _ := search.NewClgrSearcher(op)

	cs.SetUseRegularMatch(arg.Parameters["regular"].IsExist)
	cs.SetIgnoreCase(arg.Parameters["ignorecase"].IsExist)
	cs.SetMatchWholeWord(arg.Parameters["matchwholeword"].IsExist)
	cs.SetIgnoreFolderName(arg.Parameters["ignorefoldername"].IsExist)
	cs.SetDestinationDirectors(arg.ReadValue("dir"))
	cs.SetTargetKeywords(arg.ReadValue("text"))
	cs.SetDestinationFiles(arg.ReadValue("file"))
	cs.Search()
}