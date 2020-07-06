package brackets

import "errors"

type stack struct {
	s []rune
}

func newStack() *stack {
	return &stack{s: make([]rune, 0)}
}

func (s *stack) push(c rune) {
	s.s = append(s.s, c)
}

func (s *stack) pop() (rune, error) {
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("empty")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) empty() bool {
	return len(s.s) == 0
}

var openBrackets = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

var closeBrackets = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

// Bracket checks if brackets are properly nested in s
func Bracket(s string) bool {
	if s == "" {
		return true
	}
	balanced := true
	st := newStack()
	for _, c := range s {
		_, open := openBrackets[c]
		matching, close := closeBrackets[c]
		if !open && !close {
			continue
		}
		if open {
			st.push(c)
		}
		if close {
			popped, err := st.pop()
			if err != nil || matching != popped {
				return false
			}
		}
	}

	return balanced && st.empty()
}
