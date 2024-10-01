package users_domain

import (
	"fmt"

	"github.com/4strodev/du_backend/pkg/shared/domain/value_object"
	"github.com/google/uuid"
)

type UserId struct {
	value_object.ValueObject[string]
}

func NewUserId(id string) (UserId, error) {
	err := uuid.Validate(id)
	if err != nil {
		return UserId{}, fmt.Errorf("error parsing user id: %w", err)
	}

	userId := UserId{
		value_object.NewValueObject(id),
	}

	return userId, nil
}
