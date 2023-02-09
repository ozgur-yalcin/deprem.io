package models

import "encoding/json"

type Response struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (res *Response) JSON() []byte {
	json, _ := json.Marshal(res)
	return json
}
