package aenianos

import (
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model

	Nome string
	Link string
}
