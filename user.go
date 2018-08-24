package aenianos

import (
	"github.com/jinzhu/gorm"
)

// Permissions are cumulative,
// so not all admins can delete users,
// but everyone that can delete a user
// is an admin
const (
	ACCESS_NORMAL int = iota
	ACCESS_ADMIN
	ACCESS_DEL_USER
)

type User struct {
	gorm.Model

	Username string
	Password string

	AccessLevel int
}
