package shared

import (
	"reflect"
	"testing"
)

type TestCase struct {
	index int
	context interface{}
}

type Validation struct {
	T       *testing.T
	testCase TestCase
}

type ValidationConfig struct {
	Expected interface{}
	Given    interface{}
	Message  string
}

func (validation *Validation) SetTestCase(index int, context interface{}) {
	validation.testCase = TestCase {
		index: index,
		context: context,
	}
}

func (validation *Validation) IsEqual(config ValidationConfig) {
	if !reflect.DeepEqual(config.Expected, config.Given) {
		validation.equalThrow(config)
	}
}

func (validation *Validation) IsNotEqual(config ValidationConfig) {
	if reflect.DeepEqual(config.Expected, config.Given) {
		validation.equalThrow(config)
	}
}

func (validation *Validation) IsEqualError(config ValidationConfig) {
	var expectedMessage = config.Expected
	var err = config.Given
	var errorMessage = ""
	if err != nil {
		errorMessage = err.(error).Error()
	}
	validation.IsEqual(ValidationConfig{
		Expected: expectedMessage,
		Given: errorMessage,
		Message: "Not Equal Errors",
	})
}

func (validation *Validation) IsNotEmpty(config ValidationConfig) {
	if config.Expected == "" {
		validation.T.Errorf("\nExpected '%s' not to be empty", config.Message)
	}
}

func (validation *Validation) equalThrow(config ValidationConfig) {
	var message = config.Message
	if message == "" {
		message = "Not Equal Values"
	}
	if validation.testCase.context != nil {
		validation.T.Errorf(`
%s

Test case %d
Context:
%+v

Expected:	%#v
to be		%#v
`,
			message, validation.testCase.index, validation.testCase.context, config.Expected, config.Given)
	} else {
		validation.T.Errorf(`
%s
Expected:	%#v
to be		%#v
`,
			message, config.Expected, config.Given)
	}

}
