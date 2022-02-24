package questions

import (
	"strings"
)

func IsBalance(charString string) bool {
	stk := Stack{}
	for _, char := range strings.Split(charString, "") {
		if char == "(" {
			stk.Push(char)
		} else if char == ")" {
			if stk.IsEmpty() {
				return false
			}
			_, err := stk.Pop()
			if err != nil {
				panic(err)
			}
		}
	}
	return stk.IsEmpty()
}
