package rest

import (
	"reflect"
	"testing"
)

func TestRestStorageNotSingleton(t *testing.T) {
	storage1 := MakeStorage()
	storage2 := MakeStorage()
	if storage1 == storage2 {
		t.Errorf("Storage must not be singletone")
	}
}

func TestRestStorageAdd(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()

	actual := storage.Add(config)

	if actual.ID == "" {
		t.Errorf("Id must be defined")
	}
	config.ID = actual.ID
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestStorageAddOnlyOneEntry(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()

	storage.Add(config)

	size := storage.Size()
	if size != 1 {
		t.Errorf("Storage size must be %d but was %d", 1, size)
	}
}

func TestRestStorageAddWhenIdProvided(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	config.ID = "my-id"

	actual := storage.Add(config)

	if actual.ID == config.ID {
		t.Errorf("Id must be overriden")
	}
}

func TestRestStorageGet(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual := storage.Add(config)

	actualAgain, err := storage.Get(actual.ID)

	if err != nil {
		t.Errorf(err.Error())
	}
	if !reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
}

func TestRestStorageGetWhenSameId(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual := storage.Add(config)

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

func TestRestStorageGetReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeStorage()

	_, err := storage.Get("wrong-id")

	expectedError := "Rest configuration with id=wrong-id not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestRestStoragePut(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual := storage.Add(config)
	config.Request.Path = "new-path"
	config.Response.Body = "new body"

	modified, _ := storage.Put(actual.ID, config)

	if actual.ID != modified.ID {
		t.Errorf("Id must be equal:\nAdd: %s\nPut: %s", actual.ID, modified.ID)
	}
	config.ID = modified.ID
	if !reflect.DeepEqual(config, modified) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, modified)
	}
}

func TestRestStoragePutReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()

	_, err := storage.Put("wrong-id", config)

	expectedError := "Rest configuration with id=wrong-id not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestRestStorageDelete(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual := storage.Add(config)

	err := storage.Delete(actual.ID)

	if err != nil {
		t.Errorf("Must not be error, as %+v", err)
	}
	size := storage.Size()
	if size != 0 {
		t.Errorf("Storage size must be %d but was %d", 0, size)
	}
}

func TestRestStorageDeleteReturnErrorWhenGivenBadId(t *testing.T) {
	storage := MakeStorage()

	err := storage.Delete("wrong-id")

	expectedError := "Rest configuration with id=wrong-id not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error: %v\nActual error: %v", expectedError, err)
	}
}

func TestRestStorageWhenParameterChangedNoEffectOnStorage(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual := storage.Add(config)

	config.Response.Body = "new-body"

	config.ID = actual.ID
	actualAgain, _ := storage.Get(actual.ID)
	if !reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
	if reflect.DeepEqual(config, actualAgain) {
		t.Errorf("Change in config must not effect storage:\nConfig: %+v\nStorage: %+v", config, actualAgain)
	}
}

func TestRestStorageWhenReturnChangedNoEffectOnStorage(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual := storage.Add(config)

	actual.Response.Body = "new-body"

	actualAgain, _ := storage.Get(actual.ID)
	if reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage entry must be the same:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
}

func TestRestStorageGetAll(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	actual1 := storage.Add(config)
	actual2 := storage.Add(config)
	expected := []EndpointRestDto{actual1, actual2}

	configs := storage.GetAll()

	if !reflect.DeepEqual(expected, configs) {
		t.Errorf("Must be the same:\nExpected: %+v\nActual: %+v", expected, configs)
	}
}

func TestRestStorageGetAllWhenReturnChangedNoEffectOnStorage(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	storage.Add(config)
	storage.Add(config)
	actual := storage.GetAll()

	actual[0].Response.Body = "new-body"
	actual[1].Response.Body = "new-body-2"

	actualAgain := storage.GetAll()
	if reflect.DeepEqual(actual, actualAgain) {
		t.Errorf("Storage msut be immutable:\nBefore: %+v\nAfter: %+v", actual, actualAgain)
	}
}

func TestRestStorageDeleteAll(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	storage.Add(config)
	storage.Add(config)

	storage.DeleteAll()

	size := storage.Size()
	if size != 0 {
		t.Errorf("Storage must be enmpty, but had %d entries", size)
	}
}

func TestRestGlobalStorageAdd(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()

	actual := storage.AddGlobal(config)

	if actual.ID != "" {
		t.Errorf("Id must be empty")
	}
	config.ID = actual.ID
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestGlobalStorageAddOverrides(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	storage.AddGlobal(config)
	config.Request.Path = "new-path"

	actual := storage.AddGlobal(config)

	if actual.ID != "" {
		t.Errorf("Id must be empty")
	}
	config.ID = actual.ID
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestGlobalStorageGetEmpty(t *testing.T) {
	storage := MakeStorage()
	config := EndpointRestDto{}

	actual := storage.GetGlobal()

	if actual.ID != "" {
		t.Errorf("Id must be empty")
	}
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestGlobalStorageGetWhenExist(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	storage.AddGlobal(config)

	actual := storage.GetGlobal()

	if actual.ID != "" {
		t.Errorf("Id must be empty")
	}
	config.ID = actual.ID
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestGlobalStorageDelete(t *testing.T) {
	storage := MakeStorage()
	config := EndpointRestDto{}

	storage.DeleteGlobal()

	actual := storage.GetGlobal()
	if actual.ID != "" {
		t.Errorf("Id must be empty")
	}
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func TestRestGlobalStorageDeleteWhenExist(t *testing.T) {
	storage := MakeStorage()
	config := getEndpointRestDto()
	storage.AddGlobal(config)

	storage.DeleteGlobal()

	actual := storage.GetGlobal()
	config = EndpointRestDto{}
	if actual.ID != "" {
		t.Errorf("Id must be empty")
	}
	config.ID = actual.ID
	if !reflect.DeepEqual(config, actual) {
		t.Errorf("Must Be Equal:\nExpected: %+v\nActual: %+v", config, actual)
	}
}

func getEndpointRestDto() EndpointRestDto {
	return EndpointRestDto{
		ID: "",
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
			Body:     "body",
			BodyFile: "file",
			Status:   200,
			Headers: map[string]string{
				"header-response-1": "header-response-value-1",
				"header-response-2": "header-response-value-2",
			},
		},
	}
}
