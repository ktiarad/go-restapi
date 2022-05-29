package params

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status         int         `json:"status"`
	Error          string      `json:"error,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
	Message        string      `json:"message,omitempty"`
	AdditionalInfo interface{} `json:"additional_info,omitempy"`
}

func WriteJsonResponse(rw http.ResponseWriter, payload *Response) {
	rw.WriteHeader(payload.Status)
	rw.Header().Set("content-Type", "application/json")
	json.NewEncoder(rw).Encode(payload)
}
