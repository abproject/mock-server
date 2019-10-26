package rest

import (
	"github.com/abproject/mock-server/internal/rest/restentity"
	"github.com/abproject/mock-server/internal/rest/restmodels"
)

type restStorage struct {
	data     map[string]*entityRest
	global   *entityRest
	entities restentity.StorageRestEntity
}

type entityRest struct {
	Config         restmodels.EndpointRestDto
	sequenceNumber int
}
