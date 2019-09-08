package file

import (
	"fmt"
	"sort"

	"github.com/abproject/mock-server/internal/shared"
)

// StorageFile File Repository
type StorageFile interface {
	Add(name string, data []byte) File
	Get(id string) (File, error)
	Put(id string, name string, data []byte) File
	Delete(id string) error
	DeleteAll()
	GetAll() []File
	Size() int
	IsExist(id string) bool
	GetBody(id string) ([]byte, error)
}

// MakeStorage Create new Storage
func MakeStorage() StorageFile {
	return &fileStorage{
		data: make(map[string]*fileEntity),
	}
}

func (storage *fileStorage) Add(name string, data []byte) File {
	id := shared.GetRandomId()
	body := make([]byte, len(data))
	copy(body, data)
	storage.data[id] = &fileEntity{
		name: name,
		body: body,
	}
	return File{
		ID:     id,
		Name:   storage.data[id].name,
		Length: len(storage.data[id].body),
	}
}

func (storage *fileStorage) Get(id string) (File, error) {
	if entry, ok := storage.data[id]; ok {
		return File{
			ID:     id,
			Name:   storage.data[id].name,
			Length: len(entry.body),
		}, nil
	}
	return File{}, fmt.Errorf("File with id=%s not found", id)
}

func (storage *fileStorage) Put(id string, name string, data []byte) File {
	body := make([]byte, len(data))
	copy(body, data)
	storage.data[id] = &fileEntity{
		name: name,
		body: body,
	}
	return File{
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

func (storage *fileStorage) GetAll() []File {
	files := make([]File, len(storage.data))
	i := 0
	for id := range storage.data {
		files[i] = File{
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
