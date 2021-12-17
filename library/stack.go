package library

type Element interface {
}

type Stack struct {
	s []Element
}

func (s *Stack) Len() int {
	return len(s.s)

}
func (s *Stack) Push(res Element) {
	s.s = append(s.s, res)
}
func (s *Stack) Pop() Element {

	if len(s.s) == 0 {
		panic("wtf")
	}
	r := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return r
}
