package aenianos

import (
	"github.com/jinzhu/gorm"
)

type Genero struct {
	gorm.Model

	Nome string
}
