package rest

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/shared"
)

type restStorage struct {
	data     map[string]*entityRest
	global   *entityRest
	entities models.StorageRestEntity
}

type entityRest struct {
	Config         models.EndpointRestDto
	SequenceNumber int
}

var increment = 0

// MakeStorage Create new Storage
func MakeStorage() models.StorageRest {
	return &restStorage{
		data:     make(map[string]*entityRest),
		global:   &entityRest{},
		entities: MakeEntityStorage(),
	}
}

func (storage *restStorage) Add(config models.EndpointRestDto) models.EndpointRestDto {
	id := shared.GetRandomId()
	config.ID = id
	storage.data[id] = &entityRest{
		Config:         config,
		SequenceNumber: increment,
	}
	increment++
	return storage.data[id].Config
}

func (storage *restStorage) Get(id string) (models.EndpointRestDto, error) {
	if entry, ok := storage.data[id]; ok {
		return mergeConfigs(storage.global.Config, entry.Config), nil
	}
	return models.EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) Put(id string, config models.EndpointRestDto) (models.EndpointRestDto, error) {
	if _, ok := storage.data[id]; ok {
		config.ID = id
		storage.data[id] = &entityRest{
			Config: config,
		}
		return storage.data[id].Config, nil
	}
	return models.EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) Delete(id string) error {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return nil
	}
	return fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) GetAll() []models.EndpointRestDto {
	data := []*entityRest{}
	for _, value := range storage.data {
		data = append(data, value)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].SequenceNumber < data[j].SequenceNumber
	})

	configs := make([]models.EndpointRestDto, len(data))
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

func (storage *restStorage) FindByRequest(r *http.Request, c models.AppContext) (models.EndpointRestDto, error) {
	var filtered []entityRest
	for _, entity := range storage.data {
		entity.Config = mergeConfigs(storage.global.Config, entity.Config)
		if IsEqual(*entity, r) {
			filtered = append(filtered, *entity)
		}
	}
	count := len(filtered)
	if count == 0 {
		restDto, err := storage.FindEntityByRequest(r, c)
		if err == nil {
			return restDto, nil
		}
		return models.EndpointRestDto{}, errors.New("No Endpoint Found")
	} else if count > 1 {
		sort.Slice(filtered, func(i, j int) bool {
			return Compare(filtered[i].Config.Request, filtered[j].Config.Request)
		})
	}
	return filtered[0].Config, nil
}

func (storage *restStorage) AddGlobal(config models.EndpointRestDto) models.EndpointRestDto {
	config.ID = ""
	storage.global = &entityRest{
		Config:         config,
		SequenceNumber: 0,
	}
	return storage.global.Config
}

func (storage *restStorage) GetGlobal() models.EndpointRestDto {
	if storage.global != nil {
		return storage.global.Config
	}
	return models.EndpointRestDto{}
}

func (storage *restStorage) DeleteGlobal() {
	storage.global = nil
}

func mergeConfigs(global models.EndpointRestDto, endpoint models.EndpointRestDto) models.EndpointRestDto {
	endpoint.Request = mergeRequests(global.Request, endpoint.Request)
	endpoint.Response = mergeResponse(global.Response, endpoint.Response)
	return endpoint
}

func mergeRequests(global models.RequestRestDto, endpoint models.RequestRestDto) models.RequestRestDto {
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

func mergeResponse(global models.ResponseRestDto, endpoint models.ResponseRestDto) models.ResponseRestDto {
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
func (storage *restStorage) AddEntity(config models.EntityRestDto) models.EntityRestDto {
	return storage.entities.AddEntity(config)
}

func (storage *restStorage) GetEntity(id string) (models.EntityRestDto, error) {
	return storage.entities.GetEntity(id)
}

func (storage *restStorage) PutEntity(id string, config models.EntityRestDto) (models.EntityRestDto, error) {
	return storage.entities.PutEntity(id, config)
}

func (storage *restStorage) DeleteEntity(id string) error {
	return storage.entities.DeleteEntity(id)
}

func (storage *restStorage) GetAllEntities() []models.EntityRestDto {
	return storage.entities.GetAllEntities()
}

func (storage *restStorage) DeleteAllEntities() {
	storage.entities.DeleteAllEntities()
}

func (storage *restStorage) SizeEntities() int {
	return storage.entities.SizeEntities()
}

func (storage *restStorage) FindEntityByRequest(r *http.Request, c models.AppContext) (models.EndpointRestDto, error) {
	return storage.entities.FindEntityByRequest(r, c)
}
