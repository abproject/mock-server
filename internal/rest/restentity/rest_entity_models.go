package restentity

// EntityRestDto Rest Entity Data Transfer Object
type EntityRestDto struct {
	Name string `json:"name" yaml:"name"`
	File string `json:"file" yaml:"file"`
	ID   string `json:"id" yaml:"id"`
}
