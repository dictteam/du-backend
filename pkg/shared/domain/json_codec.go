package domain

type ToJson interface {
	// ToJson returns a string representating the value to its corresponding json representation
	ToJson() ([]byte, error)
}

type FromJson interface {
	// FromJson mutates the state of a value with given json
	FromJson(json []byte) error
}
