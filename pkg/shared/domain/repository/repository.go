package repository

type Repository[T any] interface {
	Save(T) error
	Find() ([]T, error)
	Search() ([]T, error)
	Delete() error
}
