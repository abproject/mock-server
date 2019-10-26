package rest

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/abproject/mock-server/internal/rest/restentity"
	"github.com/abproject/mock-server/internal/rest/restmodels"
	"github.com/abproject/mock-server/internal/shared"
)

// StorageRest Rest Repository
type StorageRest interface {
	Add(config restmodels.EndpointRestDto) restmodels.EndpointRestDto
	Get(id string) (restmodels.EndpointRestDto, error)
	Put(id string, config restmodels.EndpointRestDto) (restmodels.EndpointRestDto, error)
	Delete(id string) error
	DeleteAll()
	GetAll() []restmodels.EndpointRestDto
	FindByRequest(r *http.Request) (restmodels.EndpointRestDto, error)
	Size() int
	AddGlobal(config restmodels.EndpointRestDto) restmodels.EndpointRestDto
	GetGlobal() restmodels.EndpointRestDto
	DeleteGlobal()
	restentity.StorageRestEntity
}

var increment = 0

// MakeStorage Create new Storage
func MakeStorage() StorageRest {
	return &restStorage{
		data:     make(map[string]*entityRest),
		global:   &entityRest{},
		entities: restentity.MakeEntityStorage(),
	}
}

func (storage *restStorage) Add(config restmodels.EndpointRestDto) restmodels.EndpointRestDto {
	id := shared.GetRandomId()
	config.ID = id
	storage.data[id] = &entityRest{
		Config:         config,
		sequenceNumber: increment,
	}
	increment++
	return storage.data[id].Config
}

func (storage *restStorage) Get(id string) (restmodels.EndpointRestDto, error) {
	if entry, ok := storage.data[id]; ok {
		return mergeConfigs(storage.global.Config, entry.Config), nil
	}
	return restmodels.EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) Put(id string, config restmodels.EndpointRestDto) (restmodels.EndpointRestDto, error) {
	if _, ok := storage.data[id]; ok {
		config.ID = id
		storage.data[id] = &entityRest{
			Config: config,
		}
		return storage.data[id].Config, nil
	}
	return restmodels.EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) Delete(id string) error {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return nil
	}
	return fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) GetAll() []restmodels.EndpointRestDto {
	data := []*entityRest{}
	for _, value := range storage.data {
		data = append(data, value)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].sequenceNumber < data[j].sequenceNumber
	})

	configs := make([]restmodels.EndpointRestDto, len(data))
	i := 0
	for k := range data {
		configs[i] = mergeConfigs(storage.global.Config, data[k].Config)
		i++
	}

	return configs
}

func (storage *restStorage) DeleteAll() {
	storage.data = make(map[string]*entityRest)
}

func (storage *restStorage) Size() int {
	return len(storage.data)
}

func (storage *restStorage) FindByRequest(r *http.Request) (restmodels.EndpointRestDto, error) {
	var filtered []entityRest
	for _, entity := range storage.data {
		entity.Config = mergeConfigs(storage.global.Config, entity.Config)
		if IsEqual(*entity, r) {
			filtered = append(filtered, *entity)
		}
	}
	count := len(filtered)
	if count == 0 {
		restDto, err := storage.FindEntityByRequest(r)
		if err != nil {
			return restDto, nil
		}
		return restmodels.EndpointRestDto{}, errors.New("No Entity Found")
	} else if count > 1 {
		sort.Slice(filtered, func(i, j int) bool {
			return Compare(filtered[i].Config.Request, filtered[j].Config.Request)
		})
	}
	return filtered[0].Config, nil
}

func (storage *restStorage) AddGlobal(config restmodels.EndpointRestDto) restmodels.EndpointRestDto {
	config.ID = ""
	storage.global = &entityRest{
		Config:         config,
		sequenceNumber: 0,
	}
	return storage.global.Config
}

func (storage *restStorage) GetGlobal() restmodels.EndpointRestDto {
	if storage.global != nil {
		return storage.global.Config
	}
	return restmodels.EndpointRestDto{}
}

func (storage *restStorage) DeleteGlobal() {
	storage.global = nil
}

func mergeConfigs(global restmodels.EndpointRestDto, endpoint restmodels.EndpointRestDto) restmodels.EndpointRestDto {
	endpoint.Request = mergeRequests(global.Request, endpoint.Request)
	endpoint.Response = mergeResponse(global.Response, endpoint.Response)
	return endpoint
}

func mergeRequests(global restmodels.RequestRestDto, endpoint restmodels.RequestRestDto) restmodels.RequestRestDto {
	if endpoint.Path == "" {
		endpoint.Path = global.Path
	}
	if endpoint.PathReg == "" {
		endpoint.PathReg = global.PathReg
	}
	if endpoint.Method == "" {
		endpoint.Method = global.Method
	}
	if endpoint.Headers == nil {
		endpoint.Headers = global.Headers
	}
	return endpoint
}

func mergeResponse(global restmodels.ResponseRestDto, endpoint restmodels.ResponseRestDto) restmodels.ResponseRestDto {
	if endpoint.Body == "" {
		endpoint.Body = global.Body
	}
	if endpoint.BodyFile == "" {
		endpoint.BodyFile = global.BodyFile
	}
	if endpoint.Status == 0 {
		endpoint.Status = global.Status
	}
	if endpoint.Headers == nil {
		endpoint.Headers = global.Headers
	}
	return endpoint
}

// Entity
func (storage *restStorage) AddEntity(config restmodels.EntityRestDto) restmodels.EntityRestDto {
	return storage.entities.AddEntity(config)
}

func (storage *restStorage) GetEntity(id string) (restmodels.EntityRestDto, error) {
	return storage.entities.GetEntity(id)
}

func (storage *restStorage) PutEntity(id string, config restmodels.EntityRestDto) (restmodels.EntityRestDto, error) {
	return storage.entities.PutEntity(id, config)
}

func (storage *restStorage) DeleteEntity(id string) error {
	return storage.entities.DeleteEntity(id)
}

func (storage *restStorage) GetAllEntities() []restmodels.EntityRestDto {
	return storage.entities.GetAllEntities()
}

func (storage *restStorage) DeleteAllEntities() {
	storage.entities.DeleteAllEntities()
}

func (storage *restStorage) SizeEntities() int {
	return storage.entities.SizeEntities()
}

func (storage *restStorage) FindEntityByRequest(r *http.Request) (restmodels.EndpointRestDto, error) {
	return storage.entities.FindEntityByRequest(r)
}
