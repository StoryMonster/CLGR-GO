package output

import "../common"

type Result interface {
	AddFileSearchResult(string)
	AddTextSearchResult(string, []common.MatchedLine)
	SearchConclusion()
}
