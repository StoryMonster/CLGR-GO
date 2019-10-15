package output

import (
	"log"
	"io"
	"fmt"
	"../common"
)

const (
	Debug = 1
	Info = 2
	Warn = 3
	Error = 4
)

type Output struct {
	logger *log.Logger
	result *log.Logger
	level int
	matchedFilesNum int    // 搜索到的匹配文件数量（包含文件夹名搜索结果）
	matchedLinesNum int    // 搜索到的匹配文本行数（仅文本搜索有效）
}

func New(logWriter io.Writer, resWriter io.Writer, level int) *Output {
	logger := log.New(logWriter, "[CLGR]", log.Ldate | log.Ltime)
	result := log.New(resWriter, "", 0)
	return &Output{logger, result, level, 0, 0}
}

func (op *Output) DEBUG(str string) {
	if op.level > Debug { return }
    op.write("[DEBUG] " + str)
}

func (op *Output) INFO(str string) {
	if op.level > Info { return }
    op.write("[INFO] " + str)
}

func (op *Output) WARN(str string) {
	if op.level > Warn { return }
    op.write("[WARN] " + str)
}

func (op *Output) ERROR(str string) {
	if op.level > int(Error) { return }
    op.write("[ERROR] " + str)
}

func (op *Output) RESULT(str string) {
	// XXX 专用于打印结果
	op.result.Println(str)
}

func (op *Output) AddFileSearchResult(filepath string) {
	op.matchedFilesNum++
	// XXX 此处添加逻辑以输出文件查找结果
    op.RESULT(filepath)
}

func (op *Output) AddTextSearchResult(filename string, lines []common.MatchedLine) {
	op.matchedFilesNum++
	op.matchedLinesNum += len(lines)
	// XXX 此处用于添加自定义逻辑以输出文本搜索结果
	str := fmt.Sprintf(">>>%s:\n", filename)
	for _, line := range lines {
		str += fmt.Sprintf("%d: %s\n", line.Num, line.Line)
	}
	op.RESULT(str)
}

func (op *Output) SearchConclusion() {
	str := fmt.Sprintf("Search End! Matched files: %d", op.matchedFilesNum)
	if op.matchedLinesNum > 0 {
		str += fmt.Sprintf("  Matched lines: ", op.matchedLinesNum)
	}
	op.RESULT(str)
}

func (op *Output) write(str string) {
	op.logger.Println(str)
}
