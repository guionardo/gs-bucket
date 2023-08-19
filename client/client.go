package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/guionardo/gs-bucket/domain"
)

type GSBucketClient struct {
	baseUrl string
	client  *http.Client
}

func CreateGSBucketClient(baseUrl string) *GSBucketClient {
	return &GSBucketClient{
		baseUrl: baseUrl,
		client:  http.DefaultClient,
	}
}

func (c *GSBucketClient) SetClient(client *http.Client) *GSBucketClient {
	c.client = client
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

func (c *GSBucketClient) CreatePad(name string, ttl time.Duration, deleteAfterRead bool, content []byte) (file *domain.File, err error) {
	params := url.Values{}
	params.Add("name", name)
	params.Add("ttl", ttl.String())
	if deleteAfterRead {
		params.Add("delete-after-read", "true")
	}
	body := bytes.NewReader(content)
	var req *http.Request
	url := fmt.Sprintf("%s/pads?%s", c.baseUrl, params.Encode())

	if req, err = http.NewRequest(http.MethodPost, url, body); err != nil {
		return
	}
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
