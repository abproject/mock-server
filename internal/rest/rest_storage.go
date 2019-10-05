package rest

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/abproject/mock-server/internal/shared"
)

// StorageRest Rest Repository
type StorageRest interface {
	Add(config EndpointRestDto) EndpointRestDto
	Get(id string) (EndpointRestDto, error)
	Put(id string, config EndpointRestDto) (EndpointRestDto, error)
	Delete(id string) error
	DeleteAll()
	GetAll() []EndpointRestDto
	FindByRequest(r *http.Request) (EndpointRestDto, error)
	Size() int
	AddGlobal(config EndpointRestDto) EndpointRestDto
	GetGlobal() EndpointRestDto
	DeleteGlobal()
}

var increment = 0

// MakeStorage Create new Storage
func MakeStorage() StorageRest {
	return &restStorage{
		data:   make(map[string]*entityRest),
		global: &entityRest{},
	}
}

func (storage *restStorage) Add(config EndpointRestDto) EndpointRestDto {
	id := shared.GetRandomId()
	config.ID = id
	storage.data[id] = &entityRest{
		Config:         config,
		sequenceNumber: increment,
	}
	increment++
	return storage.data[id].Config
}

func (storage *restStorage) Get(id string) (EndpointRestDto, error) {
	if entry, ok := storage.data[id]; ok {
		return mergeConfigs(storage.global.Config, entry.Config), nil
	}
	return EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) Put(id string, config EndpointRestDto) (EndpointRestDto, error) {
	if _, ok := storage.data[id]; ok {
		config.ID = id
		storage.data[id] = &entityRest{
			Config: config,
		}
		return storage.data[id].Config, nil
	}
	return EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) Delete(id string) error {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return nil
	}
	return fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) GetAll() []EndpointRestDto {
	data := []*entityRest{}
	for _, value := range storage.data {
		data = append(data, value)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].sequenceNumber < data[j].sequenceNumber
	})

	configs := make([]EndpointRestDto, len(data))
	i := 0
	for k := range data {
		configs[i] = mergeConfigs(storage.global.Config, data[k].Config)
		i++
	}

	return configs
}

func (storage *restStorage) FindByRequest(r *http.Request) (EndpointRestDto, error) {
	var filtered []entityRest
	for _, entity := range storage.data {
		entity.Config = mergeConfigs(storage.global.Config, entity.Config)
		if IsEqual(*entity, r) {
			filtered = append(filtered, *entity)
		}
	}
	count := len(filtered)
	if count == 0 {
		return EndpointRestDto{}, errors.New("No Entity Found")
	} else if count > 1 {
		sort.Slice(filtered, func(i, j int) bool {
			return Compare(filtered[i].Config.Request, filtered[j].Config.Request)
		})
	}
	return filtered[0].Config, nil
}

func (storage *restStorage) DeleteAll() {
	storage.data = make(map[string]*entityRest)
}

func (storage *restStorage) Size() int {
	return len(storage.data)
}

func (storage *restStorage) AddGlobal(config EndpointRestDto) EndpointRestDto {
	config.ID = ""
	storage.global = &entityRest{
		Config:         config,
		sequenceNumber: 0,
	}
	return storage.global.Config
}

func (storage *restStorage) GetGlobal() EndpointRestDto {
	if storage.global != nil {
		return storage.global.Config
	}
	return EndpointRestDto{}
}

func (storage *restStorage) DeleteGlobal() {
	storage.global = nil
}

func mergeConfigs(global EndpointRestDto, endpoint EndpointRestDto) EndpointRestDto {
	endpoint.Request = mergeRequests(global.Request, endpoint.Request)
	endpoint.Response = mergeResponse(global.Response, endpoint.Response)
	return endpoint
}

func mergeRequests(global RequestRestDto, endpoint RequestRestDto) RequestRestDto {
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

func mergeResponse(global ResponseRestDto, endpoint ResponseRestDto) ResponseRestDto {
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
