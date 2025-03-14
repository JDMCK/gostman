package request

type Response struct {
	StatusCode int
	Body       string
	Headers    map[string]string
}

// NewResponse initializes a new Response
func NewResponse(statusCode int, body string, headers map[string]string) *Response {
	return &Response{
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
	}
}

// GetBody returns the response body
func (r *Response) GetBody() string {
	return r.Body
}

// GetHeaders returns the response headers as a map
func (r *Response) GetHeaders() map[string]string {
	return r.Headers
}

// GetStatusCode returns the response status code
func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

// SetHeader sets a key-value pair in the headers
func (r *Response) SetHeader(key, value string) {
	r.Headers[key] = value
}

// ClearHeaders removes all headers
func (r *Response) ClearHeaders() {
	r.Headers = make(map[string]string)
}
