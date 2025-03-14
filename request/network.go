package request

import (
	"bytes"
	"io"
	"net/http"
)

// ExecuteRequest sends the Request and returns a Response
func ExecuteRequest(req *Request) (*Response, error) {
	client := &http.Client{}

	// Construct request
	httpReq, err := http.NewRequest(string(req.Method), req.URL, bytes.NewBufferString(req.Body))
	if err != nil {
		return nil, err
	}

	// Add headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Send request
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Convert headers to map
	responseHeaders := make(map[string]string)
	for key, values := range resp.Header {
		responseHeaders[key] = values[0]
	}

	// Create Response struct
	return &Response{
		StatusCode: resp.StatusCode,
		Body:       string(bodyBytes),
		Headers:    responseHeaders,
	}, nil
}
