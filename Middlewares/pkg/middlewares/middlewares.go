package middlewares

import (
	"strings"
	spro "github.com/b-entangled/GoExamples/Middlewares/pkg/string_processor"
)

type MWare func([]string)[]string

func ToLowerMiddleware(inp []string) []string{
	var ret = make([]string, 0)
	for _, row := range inp{
		ret = append(ret, strings.ToLower(row))
	}
	return ret
}

func StringsMiddlewares(inp string, next MWare) MWare{
	ret := spro.SplitString(inp)
	var out []string
	fn := func([]string)[]string{
		next(ret)
	}
	return MWare(fn)
}