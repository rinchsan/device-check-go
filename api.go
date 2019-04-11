package devicecheck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	developmentBaseURL = "https://api.development.devicecheck.apple.com/v1"
	productionBaseURL  = "https://api.devicecheck.apple.com/v1"
)

func newBaseURL(env environment) string {
	switch env {
	case Development:
		return developmentBaseURL
	case Production:
		return productionBaseURL
	default:
		panic("no matching case")
	}
}

type api struct {
	client  *http.Client
	baseURL string
}

func newAPI(env environment) api {
	return api{
		client:  new(http.Client),
		baseURL: newBaseURL(env),
	}
}

func newAPIWithHTTPClient(client *http.Client, env environment) api {
	return api{
		client:  client,
		baseURL: newBaseURL(env),
	}
}

func (api api) do(jwt, path string, requestBody interface{}) (int, []byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(requestBody); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, api.baseURL+path, buf)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt))

	resp, err := api.client.Do(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return resp.StatusCode, responseBody, nil
}
