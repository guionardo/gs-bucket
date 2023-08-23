// For testing purposes

package client

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/guionardo/gs-bucket/backend/server"
)

var sig chan os.Signal

const (
	serverAddress = "http://localhost:8080"
	masterKey     = "master.key"
)

func startServer(t *testing.T) error {
	dataFolder := t.TempDir()
	// Create master key
	masterKeyFile := path.Join(dataFolder, masterKey)
	os.WriteFile(masterKeyFile, []byte(masterKey), 0666)

	config := &server.Config{
		RepositoryFolder: dataFolder,
		HttpPort:         8080,
	}
	server.GetAuthorization().SetMasterKeyFile(masterKeyFile)

	// Listen for syscall signals for process to interrupt/quit
	sig = make(chan os.Signal, 1)
	go func() {
		if err := server.Start(sig, config); err != nil {
			t.Errorf("Failed to start server %v", err)
		}
	}()
	timeout := time.Second * 5
	if err := waitUntilRunning(timeout); err != nil {
		t.Errorf("Server doesn't started in %s", timeout)
		return err
	}

	return nil
}

func stopServer() {
	log.Printf("Signaling server to interrupt")
	sig <- os.Interrupt

	log.Printf("Waiting server to exit")
	for server.ServerIsRunning {
		fmt.Print('.')
		time.Sleep(time.Second)
	}

	log.Printf("Server exited")

}

func waitUntilRunning(timeout time.Duration) error {
	startTime := time.Now()
	for !server.ServerIsRunning {
		if time.Since(startTime) > timeout {
			return errors.New("server startup timeout")
		}
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}
