
## 介绍
clgr是一个GO语言实现的本地快速搜索文本/文件的引擎

## 使用方式
```
> clgr -h
CLGR v1.0
======================================
--ignorecase            -ic     Ignore the alphabet case.(default: false)
--matchwholeword        -mww    Match the whole word in a line or a filename when search.(default: false)
--ignorefoldername      -ifn    Ignore the name of folder when do search for files.(default: false)
--regular               -r      Use regular expression to do match(default: false)
--text                  -t      Text search keyword, multiple text can be used.
--file                  -f      File search keyword, multiple file can be used.
--dir                   -d      Specify the search directory, multiple dir can be used.(default: .)
---------------examples-----------------
clgr -f filenamekeyword                  Search files which name contains 'filenamekeyword' under current directory
clgr -f filenamekeyword -d dirname       Search files which name contains 'filenamekeyword' under directory dirname
clgr -f filenamekeyword -r               Search files which name match regular expression 'filenamekeyword' under current directory
clgr -t textkeyword                      Search textkeyword under current directory with all text files
clgr -t textkeyword -f filenamekeyword   Search textkeyword under current directory with text files with name contains filenamekeyword
```

## 如何将clgr的代码部署到自己的项目中

```
	op := output.NewDefaultOutput() // op为实现接口output.Result的实例，DefaultOutput为clgr默认处理方式
	logger := logger.New(os.Stdout, logger.Error)
	cs, _ := search.NewClgrSearcher(logger, op)
	cs.SetUseRegularMatch(true) // 正则搜索
	cs.SetIgnoreCase(true) // 忽略大小写
	cs.SetMatchWholeWord(true) // 全字匹配
	cs.SetIgnoreFolderName(true) // 不检查文件夹名字
	cs.SetDestinationDirectors([]string{}) //设置搜索的文件夹
	cs.SetTargetKeywords([]string{}) // 设置文本搜索关键字
	cs.SetDestinationFiles([]string{}) // 设置文件搜索关键字
	cs.Search()
```

output.NewDefaultOutput()是clgr默认的结果输出方式，如果要获取输出结果做额外处理，需要在自己代码中实现output.Result接口。
```
type Result interface {
	GetAFileSearchResult(string)   // 获取到一个文件搜索的结果
	GetATextSearchResult(string, []common.MatchedLine) //获取到一个文本搜索的结果
	SearchConclusion()
}
```

