package restentity

import (
	"fmt"
	"sort"
)

type restEntityStorage struct {
	data map[string]*entityRestEntity
}

type entityRestEntity struct {
	Config         EntityRestDto
	sequenceNumber int
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
}

var increment = 0

// MakeEntityStorage Create new Storage
func MakeEntityStorage() StorageRestEntity {
	return &restEntityStorage{
		data: make(map[string]*entityRestEntity),
	}
}

func (storage *restEntityStorage) AddEntity(config EntityRestDto) EntityRestDto {
	key := config.Name
	storage.data[key] = &entityRestEntity{
		Config:         config,
		sequenceNumber: increment,
	}
	increment++
	return storage.data[key].Config
}

func (storage *restEntityStorage) GetEntity(name string) (EntityRestDto, error) {
	if entry, ok := storage.data[name]; ok {
		return entry.Config, nil
	}
	return EntityRestDto{}, fmt.Errorf("Rest Entity configuration with name=%s not found", name)
}

func (storage *restEntityStorage) PutEntity(name string, config EntityRestDto) (EntityRestDto, error) {
	if _, ok := storage.data[name]; ok {
		storage.data[name] = &entityRestEntity{
			Config: config,
		}
		return storage.data[name].Config, nil
	}
	return EntityRestDto{}, fmt.Errorf("Rest Entity configuration with name=%s not found", name)
}

func (storage *restEntityStorage) DeleteEntity(name string) error {
	if _, ok := storage.data[name]; ok {
		delete(storage.data, name)
		return nil
	}
	return fmt.Errorf("Rest Entity configuration with name=%s not found", name)
}

func (storage *restEntityStorage) GetAllEntities() []EntityRestDto {
	data := []*entityRestEntity{}
	for _, value := range storage.data {
		data = append(data, value)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].sequenceNumber < data[j].sequenceNumber
	})

	configs := make([]EntityRestDto, len(data))
	i := 0
	for k := range data {
		configs[i] = data[k].Config
		i++
	}

	return configs
}

func (storage *restEntityStorage) DeleteAllEntities() {
	storage.data = make(map[string]*entityRestEntity)
}

func (storage *restEntityStorage) SizeEntities() int {
	return len(storage.data)
}
