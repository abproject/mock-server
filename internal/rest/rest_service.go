package rest

import (
	"net/http"
	"sync"
)

var instanceService iRestService
var onceService sync.Once

type iRestService interface {
	Add(dto EndpointRestDto) (EndpointOutRestDto, error)
	Get(id string) (EndpointOutRestDto, error)
	GetAll() []EndpointOutRestDto
	Update(id string, dto EndpointRestDto) (EndpointOutRestDto, error)
	Delete(id string) error
	DeleteAll()
	FindEndpoint(r *http.Request) (responseRestParsed, error)
	setStorage(storage iRestStorage)
}

type restService struct {
	storage iRestStorage
}

func GetService() iRestService {
	onceService.Do(func() {
		instanceService = newService()
	})
	return instanceService
}

func newService() iRestService {
	return &restService{
		storage: GetStorage(),
	}
}

func (service *restService) Add(dto EndpointRestDto) (EndpointOutRestDto, error) {
	endpoint, err := service.parse(dto)
	if err != nil {
		return EndpointOutRestDto{}, err
	}
	entry := restEntry{
		config:   dto,
		endpoint: endpoint,
	}
	id, entry := service.storage.Add(entry)
	result := service.mapEntryToEndpointOut(id, entry)
	return result, nil
}

func (service *restService) Get(id string) (EndpointOutRestDto, error) {
	panic("implement me")
}

func (service *restService) GetAll() []EndpointOutRestDto {
	panic("implement me")
}

func (service *restService) Update(id string, endpoint EndpointRestDto) (EndpointOutRestDto, error) {
	panic("implement me")
}

func (service *restService) Delete(id string) error {
	panic("implement me")
}

func (service *restService) DeleteAll() {
	panic("implement me")
}

func (service *restService) FindEndpoint(r *http.Request) (responseRestParsed, error) {
	panic("implement me")
}

func (service *restService) setStorage(storage iRestStorage) {
	service.storage = storage
}

func (service *restService) parse(endpoint EndpointRestDto) (endpointRestParsed, error) {
	return endpointRestParsed{}, nil
}

func (service *restService) mapEntryToEndpointOut(id string, entry restEntry) EndpointOutRestDto {
	return EndpointOutRestDto{}
}
