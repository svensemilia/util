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

func StringifyBody(response *LambdaApiResponse, body interface{}) LambdaApiResponse {
	b, _ := json.Marshal(body)
	response.Body = string(b)
	return *response
}
