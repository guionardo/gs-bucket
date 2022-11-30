package repositories

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	File struct {
		Name       string    `json:"name"`
		Created    time.Time `json:"created"`
		ValidUntil time.Time `json:"valid_until"`
		Code       string    `json:"code"`
		MimeType   string    `json:"mime_type"`
		Content    []byte    `json:"-"`
		Size       int       `json:"size"`
	}
)

const (
	FILE_AGE_BASE = 16 // HEXADECIMAL
)

func CreateFile(name string, body []byte, mimeType string, ttl time.Duration) *File {
	file := NewFile(name, ttl, mimeType)
	file.Content = body
	file.Size = len(body)
	return file
}

func NewFile(name string, ttl time.Duration, mimeType string) *File {
	validUntil := time.Now().Add(ttl).Round(time.Second)
	return &File{
		Name:       name,
		Created:    time.Now().Round(time.Second),
		ValidUntil: validUntil,
		Code:       createRepositoryName(name, validUntil),
		MimeType:   mimeType,
	}
}

func ReadFile(filename string) (*File, error) {
	if stat, err := os.Stat(filename); err != nil || stat.IsDir() {
		return nil, fmt.Errorf("file not found %s", filename)
	}
	metaFile := fmt.Sprintf("%s.meta", filename)
	if stat, err := os.Stat(metaFile); err != nil || stat.IsDir() {
		return nil, fmt.Errorf("meta file not found %s", metaFile)
	}
	meta, err := os.ReadFile(metaFile)
	if err != nil {
		return nil, err
	}
	file := &File{}
	if err = json.Unmarshal(meta, &file); err != nil {
		return nil, err
	}
	file.Content, err = os.ReadFile(filename)
	return file, err
}

func (file *File) Save(name string) (err error) {
	if err = os.WriteFile(name, file.Content, os.ModePerm); err != nil {
		return
	}
	metaFile := fmt.Sprintf("%s.meta", name)
	meta, err := json.Marshal(file)
	if err != nil {
		return err
	}
	return os.WriteFile(metaFile, meta, os.ModePerm)
}

func (file *File) Delete(name string) (err error) {
	if err = os.Remove(name); err != nil {
		return
	}
	metaFile := fmt.Sprintf("%s.meta", name)
	return os.Remove(metaFile)
}

func createRepositoryName(name string, validUntil time.Time) string {
	fileName := base64.StdEncoding.EncodeToString([]byte(name))
	return fmt.Sprintf("%s_%s", strconv.FormatInt(validUntil.Unix(), FILE_AGE_BASE), fileName)
}

func parseRepositoryName(fileName string) (name string, validUntil time.Time, err error) {
	splitted := strings.Split(fileName, "_")
	if len(splitted) != 2 {
		err = errors.New("invalid file name")
		return
	}
	validUntilUnix, err := strconv.ParseInt(splitted[0], FILE_AGE_BASE, 64)
	if err != nil {
		return "", time.Time{}, err
	}
	bname, err := base64.StdEncoding.DecodeString(splitted[1])
	if err != nil {
		return "", time.Time{}, err
	}
	return string(bname), time.Unix(validUntilUnix, 0), nil
}
