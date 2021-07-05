package market

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	btcMarketURL = "https://api.getwexel.com/fake-btc-markets"

	InitialDate = "2016-01-01T00:00:00Z"
)

var (
	client *httpClient
)

type httpClient struct{
	baseURL string
	client  *http.Client
}

func getClient() *httpClient {
	if client == nil {
		client = &httpClient{
			baseURL: btcMarketURL,
			client:  &http.Client{
				Timeout: 5*time.Second,
			},
		}
	}
	return client
}

func (c *httpClient) request(
	method string,
	path string,
	body []byte,
	params map[string]string,
	headers map[string]string,
) (*http.Response, error) {
	url := c.baseURL + path
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return c.client.Do(request)
}

func get(
	path string,
	params map[string]string,
	headers map[string]string,
) (*http.Response, error) {
	c := getClient()

	return c.request(http.MethodGet, path, nil, params, headers)
}

func unmarshalBody(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	var resp response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Data, v)
}

func getDefaultParams(timestamp time.Time) map[string]string {
	return map[string]string{
		"timestamp": timestamp.Format(time.RFC3339),
	}
}

type response struct{
	Data json.RawMessage `json:"data"`
	Meta json.RawMessage `json:"meta"`
}
