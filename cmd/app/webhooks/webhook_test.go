package webhooks

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type webhookTestCase struct {
	tag              string
	method           string
	urlToServe       string
	urlToHit         string
	headers          map[string]string
	body             string
	function         gin.HandlerFunc
	status           int
	responseFilePath string
}

// Creates new router in testing mode
func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	return router
}

// Used to run single API test case. It makes HTTP request and returns its response
func testWebhook(router *gin.Engine, method string, urlToServe string, urlToHit string, function gin.HandlerFunc, headers map[string]string, body string) *httptest.ResponseRecorder {
	router.Handle(method, urlToServe, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlToHit, bytes.NewBufferString(body))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	router.ServeHTTP(res, req)
	return res
}

// Used to run suite (list) of test cases. It checks JSON response is same as expected data in test case file.
// All test expected test case responses are stored in `test_data/test_case_data` folder in format `<suite_name>_t<number>.json`
func runWebhookTests(t *testing.T, tests []webhookTestCase) {
	for _, test := range tests {
		router := newRouter()
		var body []byte
		if test.body != "" {
			body, _ = ioutil.ReadFile(test.body)
		}
		res := testWebhook(router, test.method, test.urlToServe, test.urlToHit, test.function, test.headers, string(body))
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.responseFilePath != "" {
			response, _ := ioutil.ReadFile(test.responseFilePath)
			assert.JSONEq(t, string(response), res.Body.String(), test.tag)
		}
	}
}
