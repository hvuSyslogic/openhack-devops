package tripsgo

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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
	for _, test := range tests {
		res := testAPI(router, test.method, test.url, test.body)
		test.actualResponse = res.Body.String()
		log.Printf(test.actualResponse)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.expectedResponse != "" {
			assert.JSONEq(t, test.expectedResponse, res.Body.String(), test.tag)
		}
	}
}
