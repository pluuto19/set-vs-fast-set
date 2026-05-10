package setvsfastset

type SetStruct[T comparable] struct {
	set map[T]struct{}
}

func NewSetStruct[T comparable](n int) *SetStruct[T] {
	return &SetStruct[T]{
		set: make(map[T]struct{}, n),
	}
}

func (s *SetStruct[T]) Add(val T) {
	s.set[val] = struct{}{}
}

func (s *SetStruct[T]) Delete(val T) {
	delete(s.set, val)
}

func (s *SetStruct[T]) Contains(val T) bool {
	_, ok := s.set[val]
	return ok
}
