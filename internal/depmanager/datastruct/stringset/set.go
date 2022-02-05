package stringset

type Set struct {
	m map[string]struct{}
}

func New() *Set {
	return &Set{m: map[string]struct{}{}}
}

func (s *Set) Add(e string) {
	s.m[e] = struct{}{}
}

func (s *Set) Has(e string) bool {
	_, ok := s.m[e]
	return ok
}
