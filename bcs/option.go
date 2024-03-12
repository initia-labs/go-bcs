package bcs

type Option[T any] []T

func Some[T any](v T) Option[T] {
	return []T{v}
}

func None[T any]() Option[T] {
	return []T{}
}

func (opt *Option[T]) IsSome() bool {
	return opt != nil && len(*opt) == 1
}

func (opt *Option[T]) IsNone() bool {
	return opt == nil || len(*opt) == 0
}
