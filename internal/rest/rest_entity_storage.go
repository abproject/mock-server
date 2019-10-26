package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/abproject/mock-server/internal/models"
)

type restEntityStorage struct {
	data map[string]*models.EntityRestEntity
}

var incrementEntity = 0

// MakeEntityStorage Create new Storage
func MakeEntityStorage() models.StorageRestEntity {
	return &restEntityStorage{
		data: make(map[string]*models.EntityRestEntity),
	}
}

func (storage *restEntityStorage) AddEntity(config models.EntityRestDto) models.EntityRestDto {
	key := config.Name
	storage.data[key] = &models.EntityRestEntity{
		Config:         config,
		SequenceNumber: incrementEntity,
	}
	incrementEntity++
	return storage.data[key].Config
}

func (storage *restEntityStorage) GetEntity(name string) (models.EntityRestDto, error) {
	if entry, ok := storage.data[name]; ok {
		return entry.Config, nil
	}
	return models.EntityRestDto{}, fmt.Errorf("Rest Entity configuration with name=%s not found", name)
}

func (storage *restEntityStorage) PutEntity(name string, config models.EntityRestDto) (models.EntityRestDto, error) {
	if _, ok := storage.data[name]; ok {
		storage.data[name] = &models.EntityRestEntity{
			Config: config,
		}
		return storage.data[name].Config, nil
	}
	return models.EntityRestDto{}, fmt.Errorf("Rest Entity configuration with name=%s not found", name)
}

func (storage *restEntityStorage) DeleteEntity(name string) error {
	if _, ok := storage.data[name]; ok {
		delete(storage.data, name)
		return nil
	}
	return fmt.Errorf("Rest Entity configuration with name=%s not found", name)
}

func (storage *restEntityStorage) GetAllEntities() []models.EntityRestDto {
	data := []*models.EntityRestEntity{}
	for _, value := range storage.data {
		data = append(data, value)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].SequenceNumber < data[j].SequenceNumber
	})

	configs := make([]models.EntityRestDto, len(data))
	i := 0
	for k := range data {
		configs[i] = data[k].Config
		i++
	}

	return configs
}

func (storage *restEntityStorage) DeleteAllEntities() {
	storage.data = make(map[string]*models.EntityRestEntity)
}

func (storage *restEntityStorage) SizeEntities() int {
	return len(storage.data)
}

func (storage *restEntityStorage) FindEntityByRequest(r *http.Request, c models.AppContext) (models.EndpointRestDto, error) {
	regexp := regexp.MustCompile(`^\/(\w+)\/?(\d+)?$`)
	groups := regexp.FindStringSubmatch(r.RequestURI)

	if len(groups) >= 1 {
		entityName := groups[1]
		id := groups[2]

		if entity, ok := storage.data[entityName]; ok {
			a, err := storage.parseEntity(strings.ToUpper(r.Method), id, entity)

			return a, err
		}
	}
	return models.EndpointRestDto{}, errors.New("Entity not found")
}

func (storage *restEntityStorage) parseEntity(method string, id string, entity *models.EntityRestEntity) (models.EndpointRestDto, error) {
	if method == "GET" {
		if id == "" {
			return models.EndpointRestDto{
				Response: models.ResponseRestDto{
					BodyFile: entity.Config.Data,
					Status:   200,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
				},
			}, nil
		}

		var data []map[string]interface{}
		json.Unmarshal([]byte(entity.Config.Data), &data)
		fmt.Printf("Birds : %+v", data)

		return models.EndpointRestDto{
			Response: models.ResponseRestDto{
				BodyFile: entity.Config.Data,
				Status:   200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		}, nil

	} else if method == "PUT" && id != "" {

	} else if method == "DELETE" {

	} else if method == "POST" && id == "" {

	}
	return models.EndpointRestDto{}, errors.New("Entity not found")
}
