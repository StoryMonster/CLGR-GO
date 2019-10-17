package main

import (
	"./search"
	"./output"
	"./args"
	"./logger"
	"os"
)

func main() {
	arg := args.ReadCmdParameters()
	if arg == nil { return }
	op := output.NewDefaultOutput()
	logger := logger.New(os.Stdout, logger.Error)
	cs, _ := search.NewClgrSearcher(logger, op)
	cs.SetFromArgs(arg)
	cs.Search()
}