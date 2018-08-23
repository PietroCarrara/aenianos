package aenianos

import (
	"github.com/jinzhu/gorm"
)

type Status struct {
	gorm.Model

	Nome string
}
