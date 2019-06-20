package crud

import (
	httpTesting "github.com/abproject/mock-server/internal/testing"
	"testing"
)

func TestRequestInit(t *testing.T) {
	tests := []httpTesting.HttpTestCase{
		{
			Type: "GET",
			Path: "users",
			ExpectedStatus: 200,
			ExpectedBody: "[{\"id\":1,\"firstname\":\"John\",\"lastname\":\"Galt\"},{\"id\":2,\"firstname\":\"Guy\",\"lastname\":\"Montag\"}]",
			ExpectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			Type: "GET",
			Path: "users/1",
			ExpectedStatus: 200,
			ExpectedBody: "{\"id\":1,\"firstname\":\"John\",\"lastname\":\"Galt\"}",
			ExpectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			Type: "GET",
			Path: "users/2",
			ExpectedStatus: 200,
			ExpectedBody: "{\"id\":1,\"firstname\":\"John\",\"lastname\":\"Galt\"}",
			ExpectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			Type: "POST",
			Path: "users",
			ExpectedStatus: 201,
			ExpectedBody: "{\"id\":1,\"firstname\":\"John\",\"lastname\":\"Galt\"}",
			ExpectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			Type: "PUT",
			Path: "users/1",
			ExpectedStatus: 200,
			ExpectedBody: "{\"id\":1,\"firstname\":\"John\",\"lastname\":\"Galt\"}",
			ExpectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			Type: "DELETE",
			Path: "users/1",
			ExpectedStatus: 200,
			ExpectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}

	httpTesting.RunCases("CRUD", "config.yml", &tests, t)
}

