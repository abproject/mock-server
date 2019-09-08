package file

// File File structure
type File struct {
	ID     string `json:"id" yaml:"id"`
	Name   string `json:"name" yaml:"name"`
	Length int    `json:"length" yaml:"length"`
}
