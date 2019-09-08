package rest

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

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
}

// MakeStorage Create new Storage
func MakeStorage() StorageRest {
	return &restStorage{
		data: make(map[string]*entityRest),
	}
}

func (storage *restStorage) Add(config EndpointRestDto) EndpointRestDto {
	id := shared.GetRandomId()
	config.ID = id
	storage.data[id] = &entityRest{
		Config:  config,
		created: time.Now().UnixNano(),
	}
	return storage.data[id].Config
}

func (storage *restStorage) Get(id string) (EndpointRestDto, error) {
	if entry, ok := storage.data[id]; ok {
		return entry.Config, nil
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
		return data[i].created < data[j].created
	})

	configs := make([]EndpointRestDto, len(data))
	i := 0
	for k := range data {
		configs[i] = data[k].Config
		i++
	}

	return configs
}

func (storage *restStorage) FindByRequest(r *http.Request) (EndpointRestDto, error) {
	var filtered []entityRest
	for _, entity := range storage.data {
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
