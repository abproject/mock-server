package rest

import "testing"

func TestRestServiceSingleton(t *testing.T) {
	service1 := GetService()
	service2 := GetService()
	if service1 != service2 {
		t.Errorf("Service must be singletone")
	}
}

func TestRestServiceNotSingleton(t *testing.T) {
	service1 := newService()
	service2 := newService()
	if service1 == service2 {
		t.Errorf("Service must not be singletone")
	}
}

func TestRestServiceAdd(t *testing.T) {
	service := GetService()
	endpoint := EndpointRestDto{
		Request: RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header-request-1": "header-request-value-1",
				"header-request-2": "header-request-value-2",
			},
		},
		Response: ResponseRestDto{
			Body: "Body",
			File: "",
			Status: 200,
			Headers: map[string]string{
				"header-response-1": "header-response-value-1",
				"header-response-2": "header-response-value-2",
			},
		},
	}

	result, err := service.Add(endpoint)

	if err != nil {
		t.Errorf(err.Error())
	}
	expected := EndpointOutRestDto{
		result.Id,
		endpoint,
	}
	t.Errorf(expected.Id)
}

func TestRestServiceAddRequestGivenPathAndPathReg(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddThrowsWhenRequestGivenNoPathAndPathReg(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddRequestGivenPathWithDynamicId(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddRequestWhenDefaultMethod(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddRequestWhenDefaultHeaders(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddResponseGivenBodyAndFile(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddResponseThrowsWhenNoBodyAndFile(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddResponseThrowsWhenNoFileFound(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddResponseGivenDefaultStatus(t *testing.T) {
	panic("TODO")
}

func TestRestServiceAddResponseGivenDefaultHeaders(t *testing.T) {
	panic("TODO")
}

func TestRestServiceGet(t *testing.T) {
	panic("TODO")
}

func TestRestServiceGetThrowsWhenGivenBadId(t *testing.T) {
	panic("TODO")
}

func TestRestServiceGetAll(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateRequestGivenPathAndPathReg(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateThrowsWhenRequestGivenNoPathAndPathReg(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateRequestGivenPathWithDynamicId(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateRequestGivenDefaultMethod(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateRequestGivenDefaultHeaders(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateResponseGivenBodyAndFile(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateResponseThrowsWhenNoBodyAndFile(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateResponseThrowsNoFileFound(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateResponseGivenDefaultStatus(t *testing.T) {
	panic("TODO")
}

func TestRestServiceUpdateResponseGivenDefaultHeaders(t *testing.T) {
	panic("TODO")
}

func TestRestServiceDelete(t *testing.T) {
	panic("TODO")
}

func TestRestServiceDeleteThrowsWhenGivenBadId(t *testing.T) {
	panic("TODO")
}

func TestRestServiceDeleteAll(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenMethod(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenMethodAll(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenPath(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenPathReg(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenPathAndPathReg(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenHeaders(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenMultipleHeaders(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointThrowsWhenNotFound(t *testing.T) {
	panic("TODO")
}

func TestRestServiceFindEndpointGivenFitMultipleEndpoint(t *testing.T) {
	panic("TODO")
}

func getEndpointDto() EndpointRestDto {
	return EndpointRestDto{
		Request: RequestRestDto{
			Method:  "GET",
			Path:    "path",
			PathReg: "",
			Headers: map[string]string{
				"header-request-1": "header-request-value-1",
				"header-request-2": "header-request-value-2",
			},
		},
		Response: ResponseRestDto{
			Body: "Body",
			File: "",
			Status: 200,
			Headers: map[string]string{
				"header-response-1": "header-response-value-1",
				"header-response-2": "header-response-value-2",
			},
		},
	}
}
