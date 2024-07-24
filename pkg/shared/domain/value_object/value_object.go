package value_object

type ValueObject[T any] struct {
	value T
}

func NewValueObject[T any](value T) ValueObject[T] {
	return ValueObject[T]{
		value: value,
	}
}

func (v *ValueObject[T]) Value() T {
	return v.value
}
