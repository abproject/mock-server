package file

import (
	"reflect"
	"sort"
	"testing"
)

func TestFileStorageNotSingleton(t *testing.T) {
	storage1 := MakeStorage()
	storage2 := MakeStorage()
	if storage1 == storage2 {
		t.Errorf("Storage must not be singletone")
	}
}

func TestFileStorageAdd(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"

	actual := storage.Add(name, data)

	if actual.ID == "" {
		t.Errorf("Path must be defined")
	}
	if actual.Name != "file-name" {
		t.Errorf("Must Be Equal Name:\nExpected: file-name\nActual: %s", actual.Name)
	}
	if actual.Length != 3 {
		t.Errorf("Must Be Equal Length:\nExpected: 3\nActual: %d", actual.Length)
	}
}

func TestFileStorageAddOnlyOneEntry(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"

	storage.Add(name, data)

	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestFileStorageGet(t *testing.T) {
	storage := MakeStorage()
	data := []byte("get")
	name := "file-name"
	actual := storage.Add(name, data)

	actualAgain, err := storage.Get(actual.ID)

	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
	if actualAgain.ID == "" {
		t.Errorf("Path must be defined")
	}
	if actualAgain.Name != "file-name" {
		t.Errorf("Must Be Equal Name:\nExpected: file-name\nActual: %s", actualAgain.Name)
	}
	if actualAgain.Length != 3 {
		t.Errorf("Must Be Equal Length:\nExpected: 3\nActual: %d", actual.Length)
	}
}

func TestFileStorageGetWhenSameId(t *testing.T) {
	storage := MakeStorage()
	data := []byte("get")
	name := "file-name"
	actual := storage.Add(name, data)

	actualAgain1, err1 := storage.Get(actual.ID)
	actualAgain2, err2 := storage.Get(actual.ID)

	if err1 != nil {
		t.Errorf(err1.Error())
	}
	if err2 != nil {
		t.Errorf(err2.Error())
	}
	if !reflect.DeepEqual(actualAgain1, actualAgain2) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actualAgain1, actualAgain2)
	}
}

func TestFileStorageGetReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeStorage()

	_, err := storage.Get("wrong-id")

	expectedError := "File with id=wrong-id not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestFileStoragePut(t *testing.T) {
	storage := MakeStorage()
	data := []byte("put")
	name := "file-name"
	id := "id"

	actual := storage.Put(id, name, data)

	if actual.ID != id {
		t.Errorf("Id must be equal:\nAdd: %s\nPut: %s", actual.ID, id)
	}
	if actual.Name != "file-name" {
		t.Errorf("Must Be Equal Name:\nExpected: file-name\nActual: %s", actual.Name)
	}
	if actual.Length != 3 {
		t.Errorf("Must Be Equal Length:\nExpected: 8\nActual: %d", actual.Length)
	}
}

func TestFileStoragePutAddsOnlyOneEntry(t *testing.T) {
	storage := MakeStorage()
	data := []byte("put")
	name := "file-name"
	id := "id"

	storage.Put(id, name, data)

	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestFileStoragePutOverridesAdd(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"
	actual := storage.Add(name, data)
	newData := []byte("put data")
	newName := "new-file-name"

	modified := storage.Put(actual.ID, newName, newData)

	if actual.ID != modified.ID {
		t.Errorf("Id must be equal:\nAdd: %s\nPut: %s", actual.ID, modified.ID)
	}
	if modified.Name != "new-file-name" {
		t.Errorf("Must Be Equal Name:\nExpected: new-file-name\nActual: %s", modified.Name)
	}
	if modified.Length != 8 {
		t.Errorf("Must Be Equal Length:\nExpected: 8\nActual: %d", modified.Length)
	}
}

func TestFileStorageDelete(t *testing.T) {
	storage := MakeStorage()
	data := []byte("delete")
	name := "file-name"
	actual := storage.Add(name, data)

	err := storage.Delete(actual.ID)

	if err != nil {
		t.Errorf("Must not be error, as %+v", err)
	}
	size := storage.Size()
	if size != 0 {
		t.Errorf("Storage size must be %d but was %d", 0, size)
	}
}

func TestFileStorageDeleteReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeStorage()

	err := storage.Delete("wrong-id")

	expectedError := "File with id=wrong-id not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestFileStorageGetAll(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"
	actual1 := storage.Add(name, data)
	actual2 := storage.Add(name, data)
	expected := []File{actual1, actual2}
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].ID < expected[j].ID
	})

	configs := storage.GetAll()

	if !reflect.DeepEqual(expected, configs) {
		t.Errorf("Must be the same:\nExpected: %+v\nActual: %+v", expected, configs)
	}
}

func TestFileStorageGetAllWhenReturnChangedNoEffectOnStorage(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"
	storage.Add(name, data)
	storage.Add(name, data)
	actual := storage.GetAll()

	actual[0].ID = "new-id-0"
	actual[0].Name = "new-name-0"
	actual[0].Length = 0
	actual[1].ID = "new-id-1"
	actual[1].Name = "new-name-1"
	actual[1].Length = 0

	actualAgain := storage.GetAll()
	if actual[0].Length == actualAgain[0].Length ||
		actual[1].Length == actualAgain[1].Length ||
		actual[0].Name == actualAgain[0].Name ||
		actual[1].Name == actualAgain[1].Name ||
		actual[0].ID == actualAgain[0].ID ||
		actual[1].ID == actualAgain[1].ID {
		t.Errorf("Storage msut be immutable:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
}

func TestFileStorageDeleteAll(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"
	storage.Add(name, data)
	storage.Add(name, data)

	storage.DeleteAll()

	size := storage.Size()
	if size != 0 {
		t.Errorf("Storage must be enmpty, but had %d entries", size)
	}
}

func TestFileStorageIsNotExist(t *testing.T) {
	storage := MakeStorage()

	isExist := storage.IsExist("id")

	if isExist != false {
		t.Errorf("File 'id' not must exist")
	}
}

func TestFileStorageIsExist(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"
	actual := storage.Add(name, data)

	isExist := storage.IsExist(actual.ID)

	if isExist != true {
		t.Errorf("File '%s' must exist", actual.ID)
	}
}

func TestFileStorageGetBody(t *testing.T) {
	storage := MakeStorage()
	data := []byte("add")
	name := "file-name"
	actual := storage.Add(name, data)

	body, err := storage.GetBody(actual.ID)

	if err != nil {
		t.Errorf("Must not be error, as %+v", err)
	}
	if !reflect.DeepEqual(data, body) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", data, body)
	}
}

func TestFileStorageGetBodyReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeStorage()

	_, err := storage.GetBody("wrong-id")

	expectedError := "File with id=wrong-id not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}
