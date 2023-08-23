package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/guionardo/gs-bucket/backend/server"
	"github.com/guionardo/gs-bucket/domain"
)

type GSBucketClient struct {
	baseUrl string
	apiKey  string
	client  *http.Client
	logger  *log.Logger
}

func CreateGSBucketClient(baseUrl string, apiKey string) *GSBucketClient {
	return &GSBucketClient{
		baseUrl: baseUrl,
		apiKey:  apiKey,
		client:  http.DefaultClient,
		logger:  log.New(io.Discard, "", log.LUTC),
	}
}

func (c *GSBucketClient) SetClient(client *http.Client) *GSBucketClient {
	c.client = client
	return c
}
func (c *GSBucketClient) SetLogger(logger *log.Logger) *GSBucketClient {
	c.logger = logger
	return c
}

func (c *GSBucketClient) httpGet(url string) (res *http.Response, err error) {
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseUrl, url), nil); err != nil {
		return
	}
	return c.client.Do(req)
}

func (c *GSBucketClient) GetPads() ([]*domain.File, error) {
	res, err := c.httpGet("pads")
	if err == nil && res.StatusCode != http.StatusOK {
		err = fmt.Errorf("server returned status %s", res.Status)
	}
	if err != nil {
		return nil, err
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	files := make([]*domain.File, 0)
	if err = json.Unmarshal(content, &files); err != nil {
		return nil, err
	}
	return files, nil
}

func (c *GSBucketClient) GetPad(code string) (file []byte, err error) {
	res, err := c.httpGet(fmt.Sprintf("%s/%s", "pads", code))
	if err == nil && res.StatusCode != http.StatusOK {
		err = fmt.Errorf("server returned status %s", res.Status)
	}
	if err != nil {
		return nil, err
	}
	return io.ReadAll(res.Body)
}

func (c *GSBucketClient) CreatePad(name string, ttl time.Duration, deleteAfterRead bool, content []byte, slug string) (file *domain.File, err error) {
	params := url.Values{}
	params.Add("name", name)
	params.Add("ttl", ttl.String())
	if deleteAfterRead {
		params.Add("delete-after-read", "true")
	}
	if len(slug) > 0 {
		params.Add("slug", slug)
	}
	body := bytes.NewReader(content)
	var req *http.Request
	url := fmt.Sprintf("%s/pads?%s", c.baseUrl, params.Encode())

	if req, err = http.NewRequest(http.MethodPost, url, body); err != nil {
		return
	}
	req.Header.Add("api-key", c.apiKey)
	req.Header.Add("Content-Length", fmt.Sprintf("%d", len(content)))

	resp, err := c.client.Do(req)
	if err == nil {
		if resp.StatusCode != http.StatusCreated {
			err = fmt.Errorf("server returned status %s", resp.Status)
		}
	}
	if err != nil {
		return
	}

	newFileContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	file = &domain.File{}
	err = json.Unmarshal(newFileContent, &file)

	return
}

func (c *GSBucketClient) DeletePad(code string) (err error) {
	url := fmt.Sprintf("%s/pads/%s", c.baseUrl, code)
	var req *http.Request
	if req, err = http.NewRequest(http.MethodDelete, url, nil); err != nil {
		return
	}
	res, err := c.client.Do(req)
	if err == nil && res.StatusCode != http.StatusAccepted {
		err = fmt.Errorf("server returned status %s", res.Status)
	}
	return
}

func (c *GSBucketClient) CreateKeyForUser(userName string) (key domain.AuthResponse, err error) {
	var req *http.Request
	if req, err = c.createRequest(http.MethodPost, c.apiKey, nil, "auth/%s", userName); err != nil {
		return
	}
	body, _, err := c.doRequest(req, http.StatusCreated)
	if err == nil {
		err = json.Unmarshal(body, &key)
	}
	return
}

func (c *GSBucketClient) GetAuthList() (auths []domain.AuthResponse, err error) {
	var req *http.Request
	if req, err = c.createRequest(http.MethodGet, c.apiKey, nil, "auth"); err != nil {
		return
	}
	body, _, err := c.doRequest(req, http.StatusOK)
	if err == nil {
		err = json.Unmarshal(body, &auths)
	}
	return

	// url := fmt.Sprintf("%s/auth", c.baseUrl)
	// var req *http.Request
	// if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
	// 	return
	// }
	// req.Header.Add("api-key", c.apiKey)
	// if res, err := c.client.Do(req); err == nil && res.StatusCode == http.StatusOK {
	// 	if content, err := io.ReadAll(res.Body); err == nil {
	// 		auths := make([]domain.AuthResponse, 0)
	// 		err = json.Unmarshal(content, &auths)
	// 		return auths, err
	// 	}
	// }
	// return
}

func (c *GSBucketClient) createRequest(method string, apiKey string, body []byte, urlFormat string, a ...any) (*http.Request, error) {
	url := fmt.Sprintf(fmt.Sprintf("%s/%s", c.baseUrl, urlFormat), a...)
	logMsg := fmt.Sprintf("[%s] %s", method, url)
	bodyReader := bytes.NewReader(body)

	if req, err := http.NewRequest(method, url, bodyReader); err == nil {
		if bodyReader != nil {
			req.Header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
		}
		if len(apiKey) > 0 {
			req.Header.Add("api-key", apiKey)
			logMsg = fmt.Sprintf("%s [api-key]", url)
		}
		if len(body) > 0 {
			logMsg = fmt.Sprintf("%s (body %s)", url, body)
		}
		c.logger.Print(logMsg)
		return req, nil
	} else {
		c.logger.Printf("failed to create request %s -> %v", logMsg, err)
		return nil, err
	}
}

func (c *GSBucketClient) doRequest(req *http.Request, successCode int) (responseBody []byte, errorResponse *server.ErrResponse, err error) {
	var res *http.Response
	if res, err = c.client.Do(req); err != nil {
		c.logger.Printf("Failed doing request %v - %v", req, err)
		return
	}
	if responseBody, err = io.ReadAll(res.Body); err == nil {
		defer res.Body.Close()
	} else {
		c.logger.Printf("Error reading response body - %v", err)
	}
	if res.StatusCode != successCode {
		if len(responseBody) > 0 {
			errResponse := server.ErrResponse{}
			if err = json.Unmarshal(responseBody, &errResponse); err == nil {
				errorResponse = &errResponse
			}
			err = fmt.Errorf("Server answerered %s", res.Status)
		}
	}
	c.logger.Printf("Response status %s - %v", res.Status, errorResponse)

	return

}
