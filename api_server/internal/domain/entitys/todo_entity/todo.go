package type_entitys

import (
	"time"
)

type Todo struct {
	id          uint // uuid.UUID = uint
	title       Title
	description Description
	limit       Limit

	is_activate bool
	update_at   time.DateTime
}
