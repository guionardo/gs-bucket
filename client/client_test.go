package client

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/guionardo/gs-bucket/domain"
)

func TestGBucketClient_All(t *testing.T) {
	if err := startServer(t); err != nil {
		return
	}
	defer stopServer()
	c := CreateGSBucketClient(serverAddress, masterKey)
	t.Run("#0 - Create and Get Pad", func(t *testing.T) {
		content := []byte("SOME CONTENT\nMULTIPLE LINES")
		file, err := c.CreatePad("test_pad.txt", 0, false, content, "test_1")
		if err != nil {
			t.Errorf("Failed on create pad - %v", err)
			return
		}

		reloadContent, err := c.GetPad(file.Slug)
		if err != nil {
			t.Errorf("Failed on get pad %s - %v", file.Slug, err)
		} else if !reflect.DeepEqual(content, reloadContent) {
			t.Errorf("Content differs -> Expected %v -> Got %v", content, reloadContent)
		}
	})

	t.Run("#1 - Get pads", func(t *testing.T) {
		pads, err := c.GetPads()
		if err != nil {
			t.Errorf("GSBucketClient.GetPads() error %v", err)
		}
		if len(pads) != 1 {
			t.Errorf("Expected one pad on server -> got %d", len(pads))
			return
		}
	})

	t.Run("#2 - Get one-time pad", func(t *testing.T) {
		// New pad
		content := []byte("READ ONLY PAD")
		file, err := c.CreatePad("test_readonly", time.Minute*10, true, content, "test_readonetime")
		if err != nil {
			t.Errorf("Failed on create pad - %v", err)
			return
		}
		pads, err := c.GetPads()
		if err != nil {
			t.Errorf("GSBucketClient.GetPads() error %v", err)
		}
		if len(pads) < 1 {
			t.Errorf("Expected more than 0 pads on server -> got %d", len(pads))
			return
		}

		// Read one time
		reloadContent, err := c.GetPad(file.Slug)
		if err != nil {
			t.Errorf("Failed on get pad %s - %v", file.Slug, err)
		} else if !reflect.DeepEqual(content, reloadContent) {
			t.Errorf("Content differs -> Expected %v -> Got %v", content, reloadContent)
		}

		// Read second time
		_, err = c.GetPad(file.Slug)
		if err == nil {
			t.Errorf("Expected pad not found %s", file.Slug)
		}
	})

}

func TestGSBucketClient_test_requests(t *testing.T) {
	if err := startServer(t); err != nil {
		return
	}
	defer stopServer()
	var err error
	t.Run("", func(t *testing.T) {
		c := CreateGSBucketClient("http://localhost:8080", "master.key").SetLogger(log.Default())
		// Create one file
		var file *domain.File
		if file, err = c.CreatePad("test.txt", 0, true, []byte("ABCD"), "slug_test"); err != nil {
			t.Errorf("GSBucketClient.CreatePad() error = %v", err)
		} else {
			t.Logf("Pad created %v", file)
		}
		var req *http.Request
		if req, err = c.createRequest(http.MethodGet, "", nil, "pads"); err != nil {
			t.Errorf("Failed to create request %v - %v", req, err)
			return
		}

		padList := make([]domain.File, 0)
		if body, errResponse, err := c.doRequest(req, http.StatusOK); err != nil {
			t.Errorf("Failed to do request %v - %v - %v", req, err, errResponse)
			return
		} else {
			json.Unmarshal(body, &padList)
		}
		if len(padList) != 1 {
			t.Errorf("Response expected one item")
		}

	})

}

func TestGSBucketClient_CreateAndGetAuths(t *testing.T) {
	if err := startServer(t); err != nil {
		return
	}
	defer stopServer()

	t.Run("", func(t *testing.T) {
		c := CreateGSBucketClient(serverAddress, masterKey)
		key, err := c.CreateKeyForUser("test_user")
		if err != nil {
			t.Errorf("CreateKeyForUser error %v", err)
			return
		}
		keys, err := c.GetAuthList()
		if err != nil {
			t.Errorf("GetAuthList error %v", err)
			return
		}
		found := false
		for _, user := range keys {
			if key.Key == user.Key && key.UserName == user.UserName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected one key for user 'test_user'. Got %v", keys)
		}

	})

}
