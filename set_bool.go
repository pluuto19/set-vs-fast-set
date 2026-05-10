package setvsbitset

type SetBool[T comparable] struct {
	set map[T]bool
}

func NewSetBool[T comparable](n int) *SetBool[T] {
	return &SetBool[T]{
		set: make(map[T]bool, n),
	}
}

// No need to check for val, since if val exists
// Add is no-op
func (s *SetBool[T]) Add(val T) {
	s.set[val] = true
}

func (s *SetBool[T]) Delete(val T) {
	delete(s.set, val)
}

func (s *SetBool[T]) Contains(val T) bool {
	_, ok := s.set[val]
	return ok
}
