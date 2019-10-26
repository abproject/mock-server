package restmodels

// EntityRestDto Rest Entity Data Transfer Object
type EntityRestDto struct {
	Name      string `json:"name" yaml:"name"`
	Data      string `json:"dataAll" yaml:"dataAll"`
	NewEntity string `json:"dataNew" yaml:"dataNew"`
	ID        string `json:"id" yaml:"id"`
}
