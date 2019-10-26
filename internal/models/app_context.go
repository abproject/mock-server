package models

import (
	"log"
)

// AppContext Router Context
type AppContext struct {
	Logger      *log.Logger
	RestStorage *StorageRest
	FileStorage *StorageFile
	Path        string
}
