package rest

import (
	"reflect"
	"testing"
)

func TestRestStorageSingleton(t *testing.T) {
	storage1 := GetStorage()
	storage2 := GetStorage()
	if storage1 != storage2 {
		t.Errorf("Storage must be singletone")
	}
}

func TestRestStorageNotSingleton(t *testing.T) {
	storage1 := newStorage()
	storage2 := newStorage()
	if storage1 == storage2 {
		t.Errorf("Storage must not be singletone")
	}
}

func TestRestStorageAdd(t *testing.T) {
	storage := newStorage()
	entity := getEntry()

	id, _ := storage.Add(entity)

	if id == "" {
		t.Errorf("Id must be defined")
	}
	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestRestStorageAddSameEntityTwice(t *testing.T) {
	storage := newStorage()
	entity := getEntry()

	id1, _ := storage.Add(entity)
	id2, _ := storage.Add(entity)

	storageEntity1, err := storage.Get(id1)
	if err != nil {
		t.Errorf(err.Error())
	}
	storageEntity2, err := storage.Get(id2)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(entity, storageEntity1) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, storageEntity1)
	}
	if !reflect.DeepEqual(entity, storageEntity2) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, storageEntity2)
	}
	if id1 == id2 {
		t.Errorf("Id must be unique, %s != %s", id1, id2)
	}
	size := storage.Size()
	if size != 2 {
		t.Errorf("Storage size must be %d but was %d", 2, size)
	}
}

func TestRestStorageAddParameterChangeNotEffectStorage(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, storageEntity := storage.Add(entity)

	entity.config.Response.Body = "new-body"

	storageEntityAgain, err := storage.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}
	if entity.config.Response.Body == storageEntity.config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", storageEntity.config.Response.Body)
	}
	if entity.config.Response.Body == storageEntityAgain.config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", storageEntityAgain.config.Response.Body)
	}
}

func TestRestStorageAddReturnChangeNotEffectStorage(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, storageEntity := storage.Add(entity)

	storageEntity.config.Response.Body = "new-body"

	storageEntityAgain, err := storage.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}
	if storageEntity.config.Response.Body == storageEntityAgain.config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", storageEntityAgain.config.Response.Body)
	}
}

func TestRestStorageGet(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, _ := storage.Add(entity)

	storageEntity, err := storage.Get(id)

	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(entity, storageEntity) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, storageEntity)
	}
	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestRestStorageGetSameId(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, _ := storage.Add(entity)

	storageEntity1, err1 := storage.Get(id)
	storageEntity2, err2 := storage.Get(id)

	if err1 != nil {
		t.Errorf(err1.Error())
	}
	if err2 != nil {
		t.Errorf(err2.Error())
	}
	if !reflect.DeepEqual(storageEntity1, storageEntity2) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, storageEntity1)
	}
	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestRestStorageGetThrowsWhenGivenBadId(t *testing.T) {
	storage := newStorage()

	_, err := storage.Get("wrong-id")

	expectedError := "restEntry 'wrong-id' not found"
	if err.Error() != expectedError {
		t.Errorf(`
Expected error: %v
Actual error:   %v`,
			expectedError, err)
	}
}

func TestRestStorageGetReturnChangeNotEffectStorage(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, _ := storage.Add(entity)
	storageEntity, err := storage.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}

	storageEntity.config.Response.Body = "new-body"

	storageEntityAgain, err := storage.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}
	if storageEntity.config.Response.Body == storageEntityAgain.config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", storageEntity.config.Response.Body)
	}
}

func TestRestStorageGetAll(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id1, _ := storage.Add(entity)
	id2, _ := storage.Add(entity)

	entities := storage.GetAll()

	if !reflect.DeepEqual(&entity, entities[id1]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, entities[id1])
	}
	if !reflect.DeepEqual(&entity, entities[id2]) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, entities[id2])
	}
	size := len(entities)
	if size != 2 {
		t.Errorf("Storage size must be %d but was %d", 2, size)
	}
}

func TestRestStorageGetAllReturnChangeNotEffectStorage(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id1, _ := storage.Add(entity)
	id2, _ := storage.Add(entity)
	entities := storage.GetAll()

	entities[id1].config.Response.Body = "new-body"
	entities[id2].config.Response.Body = "new-body-2"

	entitiesAgain := storage.GetAll()
	if entities[id1].config.Response.Body == entitiesAgain[id1].config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", entitiesAgain[id1].config.Response.Body)
	}
	if entities[id2].config.Response.Body == entitiesAgain[id2].config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", entitiesAgain[id2].config.Response.Body)
	}
}

func TestRestStorageUpdate(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, _ := storage.Add(entity)
	entity.config.Response.Body = "new-body"

	storageEntity, err := storage.Update(id, entity)

	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(entity, storageEntity) {
		t.Errorf(`
Values must be equal
Expected: %+v
Stored:	  %+v`,
			entity, storageEntity)
	}
	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestRestStorageUpdateThrowsWhenGivenWrongId(t *testing.T) {
	storage := newStorage()
	entity := getEntry()

	_, err := storage.Update("wrong-id", entity)

	expectedError := "restEntry 'wrong-id' not found"
	if err.Error() != expectedError {
		t.Errorf(`
Expected error: %v
Actual error:   %v`,
			expectedError, err)
	}
}

func TestRestStorageUpdateParameterChangeNotEffectStorage(t *testing.T) {
	storage := newStorage()
	baseEntity := getEntry()
	id, _ := storage.Add(baseEntity)
	entity := getEntry()
	entity.config.Response.Body = "new-body"
	_, err := storage.Update(id, entity)
	if err != nil {
		t.Errorf(err.Error())
	}

	entity.config.Response.Body = "some-new-body"

	storageEntity, err := storage.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}
	if entity.config.Response.Body == storageEntity.config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", storageEntity.config.Response.Body)
	}
}

func TestRestStorageUpdateReturnChangeNotEffectStorage(t *testing.T) {
	storage := newStorage()
	baseEntity := getEntry()
	id, _ := storage.Add(baseEntity)
	entity := getEntry()
	entity.config.Response.Body = "new-body"
	storageEntity, err := storage.Update(id, entity)
	if err != nil {
		t.Errorf(err.Error())
	}

	storageEntity.config.Response.Body = "some-new-body"

	storageEntityAgain, err := storage.Get(id)
	if err != nil {
		t.Errorf(err.Error())
	}
	if storageEntity.config.Response.Body == storageEntityAgain.config.Response.Body {
		t.Errorf("Parameter must not change the storage: %s", storageEntity.config.Response.Body)
	}
}

func TestRestStorageDelete(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	id, _ := storage.Add(entity)

	err := storage.Delete(id)

	if err != nil {
		t.Errorf(err.Error())
	}
	size := storage.Size()
	if size != 0 {
		t.Errorf("Storage size must be %d but was %d", 0, size)
	}
}

func TestRestStorageDeleteThrowsWhenGivenWrongId(t *testing.T) {
	storage := newStorage()

	err := storage.Delete("wrong-id")

	expectedError := "restEntry 'wrong-id' not found"
	if err.Error() != expectedError {
		t.Errorf(`
Expected error: %v
Actual error:   %v`,
			expectedError, err)
	}
}

func TestRestStorageDeleteAll(t *testing.T) {
	storage := newStorage()
	entity := getEntry()
	storage.Add(entity)

	storage.DeleteAll()

	size := storage.Size()
	if size != 0 {
		t.Errorf("Storage size must be %d but was %d", 0, size)
	}
}

func getEntry() restEntry {
	return restEntry{
		config: EndpointRestDto{
			Request: RequestRestDto{
				Method:  "method",
				Path:    "path",
				PathReg: "path-reg",
				Headers: map[string]string{
					"header-request-1": "header-request-value-1",
					"header-request-2": "header-request-value-2",
				},
			},
			Response: ResponseRestDto{
				Body:   "body",
				File:   "file",
				Status: 200,
				Headers: map[string]string{
					"header-response-1": "header-response-value-1",
					"header-response-2": "header-response-value-2",
				},
			},
		}, endpoint: endpointRestParsed{
			request: requestRestParsed{
				method:    "method",
				path:      "path",
				isPathReg: true,
				headers: map[string][]string{
					"header-request-1": {"header-request-value-1"},
					"header-request-2": {"header-request-value-2"},
				},
			},
			response: responseRestParsed{
				body:   "body",
				file:   []byte("file"),
				status: 200,
				headers: map[string]string{
					"header-response-1": "header-response-value-1",
					"header-response-2": "header-response-value-2",
				},
			},
		},
	}
}
