package rest

import (
	"reflect"
	"sort"
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
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].ID < expected[j].ID
	})

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
			Body:   "body",
			File:   "file",
			Status: 200,
			Headers: map[string]string{
				"header-response-1": "header-response-value-1",
				"header-response-2": "header-response-value-2",
			},
		},
	}
}
