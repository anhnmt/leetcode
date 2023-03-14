package structures

func NewStack(cap int) *Stack {
	return &Stack{
		s:      make([]byte, cap),
		length: 0,
	}
}

type Stack struct {
	s      []byte
	length int
}

func (s *Stack) Push(v byte) {
	s.s[s.length] = v
	s.length++
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return 0
	}

	v := s.s[s.length-1]
	s.length--
	return v
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack) Len() int {
	return s.length
}
