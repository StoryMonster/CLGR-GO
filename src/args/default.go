package args

/*
 * ReadCmdParameters <clgr特例>直接使用此接口，可以获取已经实例化的Args对象
 */
 func ReadCmdParameters() *Args {
	arg := New("CLGR", "v1.0")
	arg.AddParameter("text", "t", []string {}, "Text search keyword, multiple text can be used.")
	arg.AddParameter("file", "f", []string {}, "File search keyword, multiple file can be used.")
	arg.AddParameter("dir", "d", []string {"."}, "Specify the search directory, multiple dir can be used.")
	arg.AddParameter("ignorecase", "ic", []string {"false"}, "Ignore the alphabet case.")
	arg.AddParameter("matchwholeword", "mww", []string {"false"}, "Match the whole word in a line or a filename when search.")
	arg.AddParameter("ignorefoldername", "ifn", []string {"false"}, "Ignore the name of folder when do search for files.")
	arg.AddParameter("regular", "r", []string {"false"}, "Use regular expression to do match")
	arg.AddExample("clgr -f filenamekeyword", "Search files which name contains 'filenamekeyword' under current directory")
	arg.AddExample("clgr -f filenamekeyword -d dirname", "Search files which name contains 'filenamekeyword' under directory dirname")
	arg.AddExample("clgr -f filenamekeyword -r", "Search files which name match regular expression 'filenamekeyword' under current directory")
	arg.AddExample("clgr -t textkeyword", "Search textkeyword under current directory with all text files")
	arg.AddExample("clgr -t textkeyword -f filenamekeyword", "Search textkeyword under current directory with text files with name contains filenamekeyword")
	arg.Parse()
	for _, param := range arg.Parameters {
		if param.IsExist { return arg }
	}
	return nil
}