package output

import (
	"io"
	"os"
	"fmt"
	"time"
	"../common"
)

type DefaultOutput struct {
	writer io.Writer
	matchedFilesNum int   // 文本搜索或者文件搜索匹配到的文件或文件夹数量
	matchedLinesNum int   // 文本搜索时匹配到的文本行数
	startTime time.Time   // 用于记录程序搜索花费的时间
}

func NewDefaultOutput() *DefaultOutput {
	return &DefaultOutput{os.Stdout, 0, 0, time.Now()}
}

func (do *DefaultOutput) write(str string) {
	str += "\n"
	do.writer.Write([]byte(str))
}

func (do *DefaultOutput) GetAFileSearchResult(filepath string) {
	do.matchedFilesNum++
	do.writer.Write([]byte(filepath+"\n"))
}

func (do *DefaultOutput) GetATextSearchResult(filename string, lines []common.MatchedLine) {
	do.matchedFilesNum++
	do.matchedLinesNum += len(lines)

	str := fmt.Sprintf(">>>%s:\n", filename)
	for _, line := range lines {
		str += fmt.Sprintf("%d: %s\n", line.Num, line.Line)
	}
	do.writer.Write([]byte(str+"\n"))
}

func (do *DefaultOutput) SearchConclusion() {
	str := fmt.Sprintf("Search End! Matched files: %d", do.matchedFilesNum)
	if do.matchedLinesNum > 0 {
		str += fmt.Sprintf("  Matched lines: %d", do.matchedLinesNum)
	}
	str += "\n"
	endTime := time.Now()
	str += fmt.Sprintf("Time cost: %fs\n", endTime.Sub(do.startTime).Seconds())
	do.writer.Write([]byte(str))
}