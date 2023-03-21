package products

import (
	"database/sql"
)

type Products struct {
	ID         int
	Name       string
	Price      int
	Stock      int
	Deleted_at sql.NullTime
	Created_by int
	Updated_at string
}