package args

import (
	"fmt"
	"os"
)

type Parameter struct {
	Key string
	Abbr string
	Value []string
	DefaultVal []string
	Desc string
	IsExist bool             // 是否存在于命令行中
}

type Args struct {
	Executable string
	Parameters map[string]Parameter
	abbrs map[string]string      // key: 简写   val: 全拼
	name string
	version string
}

const (
	StringArray = 0
	Bool = 1
	Int = 2
	String = 3
)


func New(name string, version string) *Args {
	return &Args{"", make(map[string]Parameter), make(map[string]string), name, version}
}

func (args *Args) ReadValue(key string) []string {
	param, ok := args.Parameters[key]
	if !ok { panic(fmt.Sprintf("Reading an unkonwn option: %s", key)) }
	if param.IsExist {
		return param.Value
	} else {
		return param.DefaultVal
	}
}

func (args *Args) Parse() {
	key := ""
    for idx, pair := range os.Args {
		if idx == 0 {
			args.Executable = pair
		} else {
			if pair == "--help" || pair == "-h" {
				args.Help()
				return
			}
			if (pair[0] == '-' && pair[1] == '-') {
				key = pair[2:]
				if !args.isParamAdded(key) { panic(fmt.Sprintf("Unknown option: %s", pair)) }
				param := args.Parameters[key]
				param.IsExist = true
				args.Parameters[key] = param
			} else if (pair[0] == '-') {
				abbr := pair[1:]
				if !args.isAbbrUsed(abbr) { panic(fmt.Sprintf("Unknown option: %s", pair)) }
				key, _ = args.abbrs[abbr]
				param := args.Parameters[key]
				param.IsExist = true
				args.Parameters[key] = param
			} else {
				val := pair
				param := args.Parameters[key]
				param.Value = append(param.Value, val)
				args.Parameters[key] = param
			}
		}
	}
}

func (args *Args) isParamAdded(key string) bool {
	_, ok := args.Parameters[key]
	return ok
}

func (args *Args) isAbbrUsed(abbr string) bool {
	_, ok := args.abbrs[abbr]
	return ok
}

func (args *Args) AddParameter(key string, abbr string, defaultVals []string, desc string) {
	if args.isParamAdded(key) { panic(fmt.Sprintf("Option %s is added!", key))}
	if args.isAbbrUsed(abbr) { panic(fmt.Sprintf("Abbregation [%s] is duplicate used!", abbr)) }
	param := Parameter{key, abbr, []string{}, defaultVals, desc, false}
	args.Parameters[key] = param
	args.abbrs[abbr] = key
}

func (args *Args) Help() {
	str := fmt.Sprintf("%s %s\n", args.name, args.version)
	str += "==============================\n"
	for _, param := range args.Parameters {
		if len(param.DefaultVal) > 0 {
			str += fmt.Sprintf("--%-20s  -%-5s  %-1s(default: %s)\n", param.Key, param.Abbr, param.Desc, param.DefaultVal[0])
		} else {
			str += fmt.Sprintf("--%-20s  -%-5s  %-1s\n", param.Key, param.Abbr, param.Desc)
		}
	}
	fmt.Println(str)
}