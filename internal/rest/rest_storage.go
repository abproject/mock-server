package rest

import (
	"errors"
	"fmt"
	"net/http"
	"sort"

	"github.com/abproject/mock-server/internal/shared"
)

// StorageRest Rest Repository Entity
type StorageRest interface {
	Add(config EndpointRestDto) EndpointRestDto
	Get(id string) (EndpointRestDto, error)
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
		Config: config,
	}
	return storage.data[id].Config
}

func (storage *restStorage) Get(id string) (EndpointRestDto, error) {
	if entry, ok := storage.data[id]; ok {
		return entry.Config, nil
	}
	return EndpointRestDto{}, fmt.Errorf("Rest configuration with id=%s not found", id)
}

func (storage *restStorage) GetAll() []EndpointRestDto {
	configs := make([]EndpointRestDto, len(storage.data))
	i := 0
	for k := range storage.data {
		configs[i] = storage.data[k].Config
		i++
	}
	sort.Slice(configs, func(i, j int) bool {
		return configs[i].ID < configs[j].ID
	})
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

func (storage *restStorage) Size() int {
	return len(storage.data)
}
