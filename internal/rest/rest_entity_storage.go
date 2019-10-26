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
			a, err := storage.parseEntity(strings.ToUpper(r.Method), id, entity, c)

			return a, err
		}
	}
	return models.EndpointRestDto{}, errors.New("Entity not found")
}

func (storage *restEntityStorage) parseEntity(method string, id string, entity *models.EntityRestEntity, c models.AppContext) (models.EndpointRestDto, error) {
	if method == "GET" {
		if id == "" {
			return storage.getAll(entity), nil
		}
		return storage.getOne(entity, id, c)
	} else if method == "PUT" && id != "" {
		return storage.putOne(entity, id, c)
	} else if method == "DELETE" && id != "" {
		return storage.deleteOne(entity, id, c)
	} else if method == "POST" && id == "" {
		return storage.post(entity), nil
	}
	return models.EndpointRestDto{}, errors.New("Entity not found")
}

func (storage *restEntityStorage) getAll(entity *models.EntityRestEntity) models.EndpointRestDto {
	return models.EndpointRestDto{
		Response: models.ResponseRestDto{
			BodyFile: entity.Config.Data,
			Status:   200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
}

func (storage *restEntityStorage) getOne(entity *models.EntityRestEntity, id string, c models.AppContext) (models.EndpointRestDto, error) {
	body := storage.getEntityBodyByID(entity, id, c)
	if body != "" {
		return models.EndpointRestDto{
			Response: models.ResponseRestDto{
				Body:   body,
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		}, nil
	}
	return models.EndpointRestDto{}, errors.New("Endpoint not found")
}

func (storage *restEntityStorage) post(entity *models.EntityRestEntity) models.EndpointRestDto {
	return models.EndpointRestDto{
		Response: models.ResponseRestDto{
			BodyFile: entity.Config.NewEntity,
			Status:   201,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}
}

func (storage *restEntityStorage) putOne(entity *models.EntityRestEntity, id string, c models.AppContext) (models.EndpointRestDto, error) {
	body := storage.getEntityBodyByID(entity, id, c)
	if body != "" {
		return models.EndpointRestDto{
			Response: models.ResponseRestDto{
				Body:   body,
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		}, nil
	}
	return models.EndpointRestDto{}, errors.New("Endpoint not found")
}

func (storage *restEntityStorage) deleteOne(entity *models.EntityRestEntity, id string, c models.AppContext) (models.EndpointRestDto, error) {
	body := storage.getEntityBodyByID(entity, id, c)
	if body != "" {
		return models.EndpointRestDto{
			Response: models.ResponseRestDto{
				Body:   "",
				Status: 204,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
		}, nil
	}
	return models.EndpointRestDto{}, errors.New("Endpoint not found")
}

func (storage *restEntityStorage) getEntityBodyByID(entity *models.EntityRestEntity, id string, c models.AppContext) string {
	file, _ := (*c.FileStorage).GetBody(entity.Config.Data)
	var data []map[string]interface{}
	json.Unmarshal(file, &data)
	for i := 0; i < len(data); i++ {
		dataIDInterface := data[i][entity.Config.ID]
		dataID := fmt.Sprintf("%v", dataIDInterface)
		if dataID == id {
			jsonString, _ := json.Marshal(data[i])
			return string(jsonString)
		}
	}
	return ""
}
