package middlewares

import (
	"strings"
)

type MWare func(string) []string

// ToLowerMiddleware :- Split string, Convert all words to Lower Case, Join Strings and call next Middleware/ Function
func ToLowerMiddleware(next MWare) MWare {
	fn := func(inp string) []string {
		var ret = make([]string, 0)
		str := strings.Split(inp, " ")
		for _, row := range str {
			ret = append(ret, strings.ToLower(row))
		}
		return next(strings.Join(ret, " "))
	}
	return fn
}

// RemoveRecurrent :- Split string, Remove words which will appear frequently, Join Strings and call next Middleware/ Function
func RemoveRecurrent(next MWare) MWare {
	fn := func(inp string) []string {
		var ret = make([]string, 0)
		str := strings.Split(inp, " ")
		for _, row := range str {
			ok := true
			for _, key := range []string{"at", "the", "of", "and", "by", "a", "an"} {
				if row == key {
					ok = false
					break
				}
			}
			if ok {
				ret = append(ret, row)
			}

		}
		return next(strings.Join(ret, " "))
	}
	return fn
}

// ToUpperMiddleware :- Split string, Convert all words to Upper Case, Join Strings and call next Middleware/ Function
func ToUpperMiddleware(next MWare) MWare {
	fn := func(inp string) []string {
		var ret = make([]string, 0)
		str := strings.Split(inp, " ")
		for _, row := range str {
			ret = append(ret, strings.ToUpper(row))
		}
		return next(strings.Join(ret, " "))
	}
	return fn
}
