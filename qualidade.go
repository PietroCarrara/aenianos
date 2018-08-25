package aenianos

// Qualidade representa uma qualidade de vídeo na qual
// um projeto está sendo publicado (720p, 1080p...)
type Qualidade struct {
	Nome      string
	Episodios []Episodio
}
