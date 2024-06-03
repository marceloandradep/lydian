package common

type Flags[T ~int] struct {
	flags T
}

func (f *Flags[T]) Get(flag T) bool {
	return (f.flags & flag) == flag
}

func (f *Flags[T]) Set(flag T) {
	f.flags |= flag
}

func (f *Flags[T]) Unset(flag T) {
	f.flags &= ^flag
}

func (f *Flags[T]) Clear() {
	f.flags = 0
}
