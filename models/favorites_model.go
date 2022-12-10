package models

import "github.com/google/uuid"

type Favorites struct {
	UserID    uuid.UUID
	ProductID uuid.UUID
	Product   *Product
}
