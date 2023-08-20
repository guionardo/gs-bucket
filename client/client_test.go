package client

import (
	"reflect"
	"testing"
	"time"

	"github.com/guionardo/gs-bucket/domain"
)

func TestGSBucketClient_GetPads(t *testing.T) {
	if err := startServer(t); err != nil {
		return
	}
	defer stopServer()

	t.Run("Default", func(t *testing.T) {
		c := CreateGSBucketClient("http://localhost:8080", "master.key")
		pads, err := c.GetPads()
		if err != nil {
			t.Errorf("GSBucketClient.GetPads() error %v", err)
		}
		if len(pads) != 0 {
			t.Errorf("Expected empty pads on new server -> got %d", len(pads))
			return
		}
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

		// New pad
		content = []byte("READ ONLY PAD")
		file, err = c.CreatePad("test_readonly", time.Minute*10, true, content, "test_readonetime")
		if err != nil {
			t.Errorf("Failed on create pad - %v", err)
			return
		}
		pads, err = c.GetPads()
		if err != nil {
			t.Errorf("GSBucketClient.GetPads() error %v", err)
		}
		if len(pads) != 2 {
			t.Errorf("Expected 2 pads on new server -> got %d", len(pads))
			return
		}

		// Read one time
		reloadContent, err = c.GetPad(file.Slug)
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

func TestGSBucketClient_CreateAndGetPad(t *testing.T) {
	c := CreateGSBucketClient("https://gs-bucket.fly.dev", "1234")
	sample := []byte("ABCDEFGHIJKLMNO")
	var file *domain.File
	var err error
	t.Run("", func(t *testing.T) {
		if file, err = c.CreatePad("test.txt", 0, true, sample, "slug_test"); err != nil {
			t.Errorf("GSBucketClient.CreatePad() error = %v", err)
		}
		var content []byte
		if content, err = c.GetPad(file.Slug); err != nil {
			t.Errorf("GSBucketClient.GetPad(\"%s\") error = %v", file.Slug, err)
		} else if !reflect.DeepEqual(sample, content) {
			t.Errorf("Expected pad content %v -> got %v", sample, content)
		}
		if _, err = c.GetPad(file.Slug); err == nil {
			t.Errorf("GSBucketClient.GetPad(\"%s\") should result as error", file.Slug)
		}
	})

}
