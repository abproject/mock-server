package models

import (
	"net/http"
)

// EntityRestEntity EntityRestEntity
type EntityRestEntity struct {
	Config         EntityRestDto
	SequenceNumber int
}

// StorageRestEntity Rest Entity Repository
type StorageRestEntity interface {
	AddEntity(config EntityRestDto) EntityRestDto
	GetEntity(name string) (EntityRestDto, error)
	PutEntity(name string, config EntityRestDto) (EntityRestDto, error)
	DeleteEntity(name string) error
	DeleteAllEntities()
	GetAllEntities() []EntityRestDto
	SizeEntities() int
	FindEntityByRequest(r *http.Request, c AppContext) (EndpointRestDto, error)
}
