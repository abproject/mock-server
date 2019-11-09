package httphandler

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/abproject/mock-server/shared/comparator"
	"github.com/abproject/mock-server/shared/testprinter"
)

type httpHandlerTest interface {
	makeRequest() HttpResponse
	compareRequest(request *http.Request)
	compareResponse(actualResponse HttpResponse)
}

type httpHandlerTestCase struct {
	printer    testprinter.Printer
	comparator *comparator.Comparator
	request    HttpRequest
	response   HttpResponse
	router     func(w http.ResponseWriter, r *http.Request)
}

func httpHandlerTestCaseFactory(t *testing.T) func(options *httpHandlerTestCaseOptions) httpHandlerTest {
	comparator := comparator.NewComparator()

	return func(options *httpHandlerTestCaseOptions) httpHandlerTest {
		printer := testprinter.NewTestPrinter(t, options.name)
		return &httpHandlerTestCase{
			printer:    printer,
			comparator: &comparator,
			request:    options.request,
			response:   options.response,
			router:     options.router,
		}
	}
}

func (testCase *httpHandlerTestCase) makeRequest() HttpResponse {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		testCase.compareRequest(r)
		testCase.router(w, r)
	}
	return SendHttpRequest(wrapper, &testCase.request)
}

func (testCase *httpHandlerTestCase) compareRequest(request *http.Request) {
	testCase.compareRequestUrl(request)
	testCase.compareRequestMethod(request)
	testCase.compareRequestHeaders(request)
	testCase.compareRequestBody(request)
}

func (testCase *httpHandlerTestCase) compareResponse(actualResponse HttpResponse) {
	testCase.compareResponseStatus(actualResponse)
	testCase.compareResponseHeaders(actualResponse)
	testCase.compareResponseBody(actualResponse)
}

func (testCase *httpHandlerTestCase) compareRequestUrl(request *http.Request) {
	testCase.compareEqual(strings.ToLower(testCase.request.URL), strings.ToLower(request.RequestURI), "Wrong Request URL")
}

func (testCase *httpHandlerTestCase) compareRequestMethod(request *http.Request) {
	testCase.compareEqual(strings.ToUpper(testCase.request.Method), strings.ToUpper(request.Method), "Wrong Request Method")
}

func (testCase *httpHandlerTestCase) compareRequestHeaders(request *http.Request) {
	normilizedHeaders := normilizeHeaders(request.Header)
	testCase.compareEqual(testCase.request.Headers, normilizedHeaders, "Wrong Request Headers")
}

func (testCase *httpHandlerTestCase) compareRequestBody(request *http.Request) {
	var expectedRequestBody = []byte("")
	if testCase.request.Body != nil {
		expectedRequestBody = testCase.request.Body
	}
	requestBody, _ := ioutil.ReadAll(request.Body)
	testCase.compareEqual(expectedRequestBody, requestBody, "Wrong Request Body")
}

func (testCase *httpHandlerTestCase) compareResponseStatus(actualResponse HttpResponse) {
	testCase.compareEqual(testCase.response.Status, actualResponse.Status, "Wrong Response Status")
}

func (testCase *httpHandlerTestCase) compareResponseHeaders(actualResponse HttpResponse) {
	testCase.compareEqual(testCase.response.Headers, actualResponse.Headers, "Wrong Response Headers")
}

func (testCase *httpHandlerTestCase) compareResponseBody(actualResponse HttpResponse) {
	var expectedResponseBody = []byte("")
	if testCase.response.Body != nil {
		expectedResponseBody = testCase.response.Body
	}
	testCase.compareEqual(expectedResponseBody, actualResponse.Body, "Wrong Response Body")
}

func (testCase *httpHandlerTestCase) compareEqual(expected interface{}, actual interface{}, message string) {
	if !(*testCase.comparator).Equal(expected, actual) {
		testCase.printer.ComparationError(expected, actual, message)
	}
}

func normilizeHeaders(headers map[string][]string) map[string]string {
	lowerHeaders := make(map[string]string)
	for key, values := range headers {
		normilizedKey := strings.ToLower(key)
		normilizedHeader := strings.Join(values[:], ";")
		lowerHeaders[normilizedKey] = normilizedHeader
	}
	return lowerHeaders
}
