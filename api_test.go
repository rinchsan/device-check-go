package devicecheck

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newBaseURL(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(developmentBaseURL, newBaseURL(Development))
	assert.Equal(productionBaseURL, newBaseURL(Production))
	assert.Panics(func() {
		newBaseURL(100)
	})
}

func Test_newAPI(t *testing.T) {
	env := Development
	baseURL := developmentBaseURL

	api := newAPI(env)

	assert := assert.New(t)
	assert.Equal(baseURL, api.baseURL)
	assert.NotNil(api.client)
}

func Test_newAPIWithHTTPClient(t *testing.T) {
	env := Production
	baseURL := productionBaseURL
	httpClient := new(http.Client)

	api := newAPIWithHTTPClient(httpClient, env)

	assert := assert.New(t)
	assert.Equal(baseURL, api.baseURL)
	assert.Equal(httpClient, api.client)
}

func TestAPI_do(t *testing.T) {
	api := api{
		client:  new(http.Client),
		baseURL: "http://example.com",
	}

	code, body, err := api.do("jwt", "/", nil)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, code)
	assert.NotNil(body)
	assert.Nil(err)
}

func TestAPI_do_InvalidURL(t *testing.T) {
	api := api{
		client:  new(http.Client),
		baseURL: "invalid url",
	}

	code, body, err := api.do("jwt", "/", nil)

	assert := assert.New(t)
	assert.Equal(http.StatusInternalServerError, code)
	assert.Nil(body)
	assert.NotNil(err)
}
