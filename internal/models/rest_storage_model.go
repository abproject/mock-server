package models

import (
	"net/http"
)

// StorageRest Rest Repository
type StorageRest interface {
	Add(config EndpointRestDto) EndpointRestDto
	Get(id string) (EndpointRestDto, error)
	Put(id string, config EndpointRestDto) (EndpointRestDto, error)
	Delete(id string) error
	DeleteAll()
	GetAll() []EndpointRestDto
	FindByRequest(r *http.Request, c AppContext) (EndpointRestDto, error)
	Size() int
	AddGlobal(config EndpointRestDto) EndpointRestDto
	GetGlobal() EndpointRestDto
	DeleteGlobal()
	StorageRestEntity
}
