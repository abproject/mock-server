package rest


type iRestParser interface {
	ParseRestConfig(config ConfigRestDto)
}