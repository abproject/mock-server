package models

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
