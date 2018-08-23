package aenianos

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Projeto struct {
	gorm.Model

	Nome    string
	NomeJap string
	Tipo
	Status
	Generos []Genero `gorm:"many2many:projeto_genero;"`
	Links   []Link
	Sinopse string `gorm:"type:TEXT;"`

	// Data de transmissão
	IniTrans *time.Time
	FimTrans *time.Time

	// Episódios totais e completos
	EpiTotal *int
	EpiCompl *int
}
