package rest

import "github.com/abproject/mock-server/internal/rest/restentity"

type restStorage struct {
	data     map[string]*entityRest
	global   *entityRest
	entities restentity.StorageRestEntity
}

type entityRest struct {
	Config         EndpointRestDto
	sequenceNumber int
}
