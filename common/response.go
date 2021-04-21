package common

// APIResponseBody type is a map-like response body for any API request
type APIResponseBody map[string]interface{}

// APIResponseCode type is to define the code of the API resonse
type APIResponseCode uint32

// API response code
const APISuccess APIResponseCode = 1

const (
	APIOMGeneralFail APIResponseCode = iota + 1000 //1000
)

// APIResponse defines the HTTP API standard response standard
type APIResponse interface {
	SetCode(code APIResponseCode)
	SetMsg(msg string)
	SetBody(body APIResponseBody)
	Build() response
}

type response struct {
	Code    APIResponseCode `json:"code"`
	Message string          `json:"message"`
	Body    APIResponseBody `json:"body"`
}

// SetCode sets the response code
func (res *response) SetCode(code APIResponseCode) {
	res.Code = code
}

// SetMsg sets the response message
func (res *response) SetMsg(msg string) {
	res.Message = msg
}

// SetMsg sets the response body
func (res *response) SetBody(body APIResponseBody) {
	res.Body = body
}

// Build returns the response
func (res *response) Build() response {
	return *res
}

// NewReponse creates default response structure
func NewReponse() APIResponse {
	res := response{
		Code:    APIOMGeneralFail,
		Message: "",
		Body:    APIResponseBody{"data": []string{}},
	}
	return &res
}
