package botutils

import (
	"fmt"
)

// RoleError is an error type for when there is a problem giving a role to a user.
type RoleError struct {
	Err error
}

func (e RoleError) Error() string {
	return fmt.Sprintf("error giving role to user: %v", e.Err)
}
