package market

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"tars/pkg/config"
)

var (
	client *httpClient
)

type httpClient struct{
	baseURL string
	client  *http.Client
}

func defaultClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Second}
}

// only for testing locally
func insecureClient() *http.Client {
	c := defaultClient()
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	c.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	return c
}

func getClient() *httpClient {
	if client == nil {
		client = &httpClient{
			baseURL: config.Get().MarketBaseURL,
			client:  insecureClient(),
		}
	}
	return client
}

func request(
	method string,
	path string,
	body []byte,
	params map[string]string,
	headers map[string]string,
) (*http.Response, error) {
	c := getClient()
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
	return request(http.MethodGet, path, nil, params, headers)
}

func post(
	path string,
	body []byte,
	params map[string]string,
	headers map[string]string,
) (*http.Response, error) {
	return request(http.MethodPost, path, body, params, headers)
}

func unmarshalBody(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
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
