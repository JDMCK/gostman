package request

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

type Request struct {
	Method  HTTPMethod
	URL     string
	Headers map[string]string
	Body    string
}

// NewRequest initializes a new Request
func NewRequest(method HTTPMethod, URL string, headers map[string]string, body string) *Request {
	return &Request{
		Method:  method,
		URL:     URL,
		Headers: headers,
		Body:    body,
	}
}

// GetURL returns the URL including query parameters if any
func (r *Request) GetURL() string {
	return r.URL
}

// GetBody returns the request body
func (r *Request) GetBody() string {
	return r.Body
}

// GetHeaders returns the headers as a map
func (r *Request) GetHeaders() map[string]string {
	return r.Headers
}

// SetHeader sets a key-value pair in the headers
func (r *Request) SetHeader(key, value string) {
	r.Headers[key] = value
}

// SetBody sets the request body
func (r *Request) SetBody(body string) {
	r.Body = body
}

// ClearHeaders removes all headers
func (r *Request) ClearHeaders() {
	r.Headers = make(map[string]string)
}
