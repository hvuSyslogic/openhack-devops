package tripsgo

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure-Samples/openhack-devops/src/MobileAppServiceV2/TripService/tripsgo"

	tripSvc "github.com/Azure-Samples/openhack-devops/src/MobileAppServiceV2/TripService/tripsgo"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type apiTestCase struct {
	tag              string
	method           string
	url              string
	body             string
	status           int
	expectedResponse string
	actualResponse   string
}

func newRouter() *mux.Router {
	router := tripSvc.NewRouter()
	return router
}

func testAPI(router *mux.Router, method, URL, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, URL, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, router *mux.Router, tests []apiTestCase) {
	for i := 0; i < len(tests); i += 1 {
		res := testAPI(router, tests[i].method, tests[i].url, tests[i].body)
		tests[i].actualResponse = res.Body.String()
		tripsgo.LogToConsole(tests[i].actualResponse)
		assert.Equal(t, tests[i].status, res.Code, tests[i].tag)
		if tests[i].expectedResponse != "" {
			assert.JSONEq(t, tests[i].expectedResponse, res.Body.String(), tests[i].tag)
		}
	}
}
