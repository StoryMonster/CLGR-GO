package output

import "../common"

type Result interface {
	GetAFileSearchResult(string)
	GetATextSearchResult(string, []common.MatchedLine)
	SearchConclusion()
}
