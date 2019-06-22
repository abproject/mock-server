package rest

import (
	"fmt"
	"github.com/abproject/mock-server/internal/shared"
	"github.com/jinzhu/copier"
	"sync"
)

var instanceStorage iRestStorage
var onceStorage sync.Once

type iRestStorage interface {
	Get(id string) (restEntry, error)
	GetAll() map[string]*restEntry
	Add(entity restEntry) (string, restEntry)
	Update(id string, entity restEntry) (restEntry, error)
	Delete(id string) error
	DeleteAll()
	Size() int
}

type restStorage struct {
	data map[string]*restEntry
}

func GetStorage() iRestStorage {
	onceStorage.Do(func() {
		instanceStorage = newStorage()
	})
	return instanceStorage
}

func newStorage() iRestStorage {
	return &restStorage{
		data: make(map[string]*restEntry),
	}
}

func (storage *restStorage) Get(id string) (restEntry, error) {
	if entity, ok := storage.data[id]; ok {
		copyEntity := restEntry{}
		copier.Copy(&copyEntity, &entity)
		return copyEntity, nil
	}
	return restEntry{}, fmt.Errorf("restEntry '%s' not found", id)
}

func (storage *restStorage) GetAll() map[string]*restEntry {
	copyMap := make(map[string]*restEntry)
	for k, v := range storage.data {
		copyEntity := restEntry{}
		copier.Copy(&copyEntity, &v)
		copyMap[k] = &copyEntity
	}
	return copyMap
}

func (storage *restStorage) Add(entity restEntry) (string, restEntry) {
	id := shared.GetRandomId()
	storage.data[id] = &entity
	return id, entity
}

func (storage *restStorage) Update(id string, newEntity restEntry) (restEntry, error) {
	if _, ok := storage.data[id]; ok {
		storage.data[id] = &newEntity
		return *storage.data[id], nil
	}
	return restEntry{}, fmt.Errorf("restEntry '%s' not found", id)
}

func (storage *restStorage) Delete(id string) error {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return nil
	}
	return fmt.Errorf("restEntry '%s' not found", id)
}

func (storage *restStorage) DeleteAll() {
	storage.data = make(map[string]*restEntry)
}

func (storage *restStorage) Size() int {
	return len(storage.data)
}
