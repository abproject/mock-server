package rest

import (
	"reflect"
	"testing"

	"github.com/abproject/mock-server/internal/models"
)

func TestRestEntityStorageNotSingleton(t *testing.T) {
	storage1 := MakeEntityStorage()
	storage2 := MakeEntityStorage()
	if storage1 == storage2 {
		t.Errorf("Storage must not be singletone")
	}
}

func TestRestEntityStorageAdd(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()

	actual := storage.AddEntity(config)

	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestEntityStorageAddOnlyOneEntry(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()

	storage.AddEntity(config)

	size := storage.SizeEntities()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestRestEntityStorageGet(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()
	actual := storage.AddEntity(config)

	actualAgain, err := storage.GetEntity(actual.Name)

	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
}

func TestRestEntityStorageGetWhenSameId(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()
	actual := storage.AddEntity(config)

	actualAgain1, err1 := storage.GetEntity(actual.Name)
	actualAgain2, err2 := storage.GetEntity(actual.Name)

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

func TestRestEntityStorageGetReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeEntityStorage()

	_, err := storage.GetEntity("wrong-name")

	expectedError := "Rest Entity configuration with name=wrong-name not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestRestEntityStoragePut(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()
	actual := storage.AddEntity(config)
	config.Data = "new-data.json"

	modified, _ := storage.PutEntity(actual.Name, config)

	if !reflect.DeepEqual(config, modified) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, modified)
	}
}

func TestRestEntityStoragePutReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()

	_, err := storage.PutEntity("wrong-name", config)

	expectedError := "Rest Entity configuration with name=wrong-name not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestRestEntityStorageDelete(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()
	actual := storage.AddEntity(config)

	err := storage.DeleteEntity(actual.Name)

	if err != nil {
		t.Errorf("Must not be error, as %+v", err)
	}
	size := storage.SizeEntities()
	if size != 0 {
		t.Errorf("Storage size must be %d but was %d", 0, size)
	}
}

func TestRestEntityStorageDeleteReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeEntityStorage()

	err := storage.DeleteEntity("wrong-name")

	expectedError := "Rest Entity configuration with name=wrong-name not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestRestEntityStorageWhenParameterChangedNoEffectOnStorage(t *testing.T) {
	storage := MakeEntityStorage()
	config := getEndpointRestEntityDto()
	actual := storage.AddEntity(config)

	config.Data = "new-data.json"

	actualAgain, _ := storage.GetEntity(actual.Name)
	if !reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
	if reflect.DeepEqual(config, actualAgain) {
		t.Errorf("Change in config must not effect storage:\nConfig: %+v\nStorage: %+v", config, actualAgain)
	}
}

func TestRestEntityStorageGetAll(t *testing.T) {
	storage := MakeEntityStorage()
	config1 := getEndpointRestEntityDto()
	config2 := getEndpointRestEntityDto()
	config2.Name = "name-2"
	actual1 := storage.AddEntity(config1)
	actual2 := storage.AddEntity(config2)
	expected := []models.EntityRestDto{actual1, actual2}

	configs := storage.GetAllEntities()

	if !reflect.DeepEqual(expected, configs) {
		t.Errorf("Must be the same:\nExpected: %+v\nActual: %+v", expected, configs)
	}
}

func TestRestEntityStorageGetAllWhenReturnChangedNoEffectOnStorage(t *testing.T) {
	storage := MakeEntityStorage()
	config1 := getEndpointRestEntityDto()
	config2 := getEndpointRestEntityDto()
	config2.Name = "name-2"
	storage.AddEntity(config1)
	storage.AddEntity(config2)
	actual := storage.GetAllEntities()

	actual[0].Data = "new-data-1.json"
	actual[1].Data = "new-data-2.json"

	actualAgain := storage.GetAllEntities()
	if reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage msut be immutable:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
}

func TestRestEntityStorageDeleteAll(t *testing.T) {
	storage := MakeEntityStorage()
	config1 := getEndpointRestEntityDto()
	config2 := getEndpointRestEntityDto()
	config2.Name = "name-2"
	storage.AddEntity(config1)
	storage.AddEntity(config2)

	storage.DeleteAllEntities()

	size := storage.SizeEntities()
	if size != 0 {
		t.Errorf("Storage must be enmpty, but had %d entries", size)
	}
}

func getEndpointRestEntityDto() models.EntityRestDto {
	return models.EntityRestDto{
		Name:      "name",
		Data:      "data.json",
		NewEntity: "data-new.json",
		ID:        "id",
	}
}
