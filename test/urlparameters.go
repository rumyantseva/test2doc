package test

import "github.com/rumyantseva/test2doc/doc/parse"

func RegisterURLVarExtractor(fn parse.URLVarExtractor) {
	parse.SetURLVarExtractor(&fn)
}
