package util

import (
	"encoding/json"
)

type LambdaApiResponse struct {
	Code    int               `json:"statusCode"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
	Base64  bool              `json:"isBase64Encoded"`
}

type LambdaApiRequest struct {
	Resource   string            `json:"resource"`
	Path       string            `json:"path"`
	Method     string            `json:"httpMethod"`
	Headers    map[string]string `json:"headers"`
	PathParams map[string]string `json:"pathParameters"`
	Body       string            `json:"body"`
	Base64     bool              `json:"isBase64Encoded"`
}

type LambdaApiBody struct {
	Error string      `json:"error,omitempty"`
	state interface{} `json:"state"`
}

func StringifyBody(response *LambdaApiResponse, body interface{}) LambdaApiResponse {
	b, _ := json.Marshal(body)
	response.Body = string(b)
	return *response
}
