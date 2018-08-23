package aenianos

import (
	"github.com/jinzhu/gorm"
)

type Tipo struct {
	gorm.Model

	Nome string
}
