package file

import (
	"fmt"
	"sort"

	"github.com/abproject/mock-server/internal/models"
	"github.com/abproject/mock-server/internal/shared"
)

type fileEntity struct {
	name string
	body []byte
}

type fileStorage struct {
	data map[string]*fileEntity
}

// MakeStorage Create new Storage
func MakeStorage() models.StorageFile {
	return &fileStorage{
		data: make(map[string]*fileEntity),
	}
}

func (storage *fileStorage) Add(name string, data []byte) models.File {
	id := shared.GetRandomId()
	body := make([]byte, len(data))
	copy(body, data)
	storage.data[id] = &fileEntity{
		name: name,
		body: body,
	}
	return models.File{
		ID:     id,
		Name:   storage.data[id].name,
		Length: len(storage.data[id].body),
	}
}

func (storage *fileStorage) Get(id string) (models.File, error) {
	if entry, ok := storage.data[id]; ok {
		return models.File{
			ID:     id,
			Name:   storage.data[id].name,
			Length: len(entry.body),
		}, nil
	}
	return models.File{}, fmt.Errorf("File with id=%s not found", id)
}

func (storage *fileStorage) Put(id string, name string, data []byte) models.File {
	body := make([]byte, len(data))
	copy(body, data)
	storage.data[id] = &fileEntity{
		name: name,
		body: body,
	}
	return models.File{
		ID:     id,
		Name:   storage.data[id].name,
		Length: len(storage.data[id].body),
	}
}

func (storage *fileStorage) Delete(id string) error {
	if _, ok := storage.data[id]; ok {
		delete(storage.data, id)
		return nil
	}
	return fmt.Errorf("File with id=%s not found", id)
}

func (storage *fileStorage) GetAll() []models.File {
	files := make([]models.File, len(storage.data))
	i := 0
	for id := range storage.data {
		files[i] = models.File{
			ID:     id,
			Name:   storage.data[id].name,
			Length: len(storage.data[id].body),
		}
		i++
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].ID < files[j].ID
	})
	return files
}

func (storage *fileStorage) DeleteAll() {
	storage.data = make(map[string]*fileEntity)
}

func (storage *fileStorage) Size() int {
	return len(storage.data)
}

func (storage *fileStorage) IsExist(id string) bool {
	_, exist := storage.data[id]
	return exist
}

func (storage *fileStorage) GetBody(id string) ([]byte, error) {
	if _, ok := storage.data[id]; ok {
		body := make([]byte, len(storage.data[id].body))
		copy(body, storage.data[id].body)
		return body, nil
	}
	return nil, fmt.Errorf("File with id=%s not found", id)
}
