package questions

import (
	"strings"
)

func IsBalance(charString string) bool {
	stk := Stack{}
	for _, char := range strings.Split(charString, "") {
		if char == "(" || char == "{" || char == "[" || char == "<" {
			stk.Push(char)
		} else if char == ")" || char == "}" || char == "]" || char == ">" {
			if stk.IsEmpty() {
				return false
			}
			value, err := stk.Pop()
			if err != nil {
				panic(err)
			}
			switch value {
			case "(":
				if char != ")" {
					return false
				}
			case "{":
				if char != "}" {
					return false
				}
			case "[":
				if char != "]" {
					return false
				}
			case "<":
				if char != ">" {
					return false
				}
			}
		}
	}
	return stk.IsEmpty()
}
