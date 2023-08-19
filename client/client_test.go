package client

import (
	"reflect"
	"testing"

	"github.com/guionardo/gs-bucket/domain"
)

func TestGSBucketClient_GetPads(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		c := CreateGSBucketClient("https://gs-bucket.fly.dev")
		pads, err := c.GetPads()
		if err != nil {
			t.Errorf("GSBucketClient.GetPads() error %v", err)
		} else {
			t.Logf("Pads: %v", pads)
		}

	})

}

func TestGSBucketClient_CreateAndGetPad(t *testing.T) {
	c := CreateGSBucketClient("https://gs-bucket.fly.dev")
	sample := []byte("ABCDEFGHIJKLMNO")
	var file *domain.File
	var err error
	t.Run("", func(t *testing.T) {
		if file, err = c.CreatePad("test.txt", 0, true, sample); err != nil {
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
