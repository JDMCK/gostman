package request

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const dir = ".gostman"
const requestFile = "requests.gost"

func ensureDir() error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}

// SerializeRequest writes a Request object into a JSON array file
func SerializeRequest(req *Request) error {
	if err := ensureDir(); err != nil {
		return err
	}
	filePath := filepath.Join(dir, requestFile)

	var requests []Request
	if data, err := os.ReadFile(filePath); err == nil {
		json.Unmarshal(data, &requests)
	}

	requests = append(requests, *req)
	data, err := json.MarshalIndent(requests, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

// DeserializeRequest reads all Request objects from a file
func DeserializeRequest() ([]Request, error) {
	filePath := filepath.Join(dir, requestFile)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var requests []Request
	if err := json.Unmarshal(data, &requests); err != nil {
		return nil, err
	}
	return requests, nil
}
