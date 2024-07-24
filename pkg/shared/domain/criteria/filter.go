package criteria

import (
	"fmt"

	"github.com/4strodev/du_backend/pkg/shared/domain/value_object"
)

type operators string

const (
	NONE               operators = "none"
	EQUAL                        = "eq"
	GREATER_THAN                 = "gt"
	LOWER_THAN                   = "lt"
	EQUAL_GREATER_THAN           = "egt"
	EQUAL_LOWER_THAN             = "elt"
	IN                           = "in"
	LIKE                         = "like"
)

type Operator struct {
	value_object.ValueObject[operators]
}

func NewOperator(operator operators) Operator {
	return Operator{
		value_object.NewValueObject(operator),
	}
}

func ParseOperator(unknownOperator string) (Operator, error) {
	var rawOperator operators
	switch operators(unknownOperator) {
	case EQUAL:
		rawOperator = EQUAL
	case GREATER_THAN:
		rawOperator = GREATER_THAN
	case LOWER_THAN:
		rawOperator = LOWER_THAN
	case EQUAL_GREATER_THAN:
		rawOperator = EQUAL_GREATER_THAN
	case EQUAL_LOWER_THAN:
		rawOperator = EQUAL_LOWER_THAN
	case IN:
		rawOperator = IN
	case LIKE:
		rawOperator = LIKE
	default:
		return NewOperator(NONE), fmt.Errorf("invalid operator: %s", unknownOperator)
	}
	return NewOperator(rawOperator), nil
}

type Field struct {
	value_object.ValueObject[string]
}

func NewField(value string) (Field, error) {
	if value == "" {
		return Field{}, fmt.Errorf("invalid field: %s", value)
	}
	return Field{
		value_object.NewValueObject(value),
	}, nil
}

type FieldValue struct {
	value_object.ValueObject[any]
}

func NewFieldValue(value any) (FieldValue, error) {
	return FieldValue{
		value_object.NewValueObject(value),
	}, nil
}

type Filter struct {
	Field      Field
	Operator   Operator
	FieldValue FieldValue
}
