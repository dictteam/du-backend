package value_object

type ValueObject[T any] struct {
	value T
}

func NewValueObject[T any](value T) ValueObject[T] {
	v := ValueObject[T]{
		value: value,
	}

	return v
}

func (v *ValueObject[T]) Value() T {
	return v.value
}
