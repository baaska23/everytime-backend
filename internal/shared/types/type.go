package types

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Length  uint   `json:"int"`
}

func NewAPIResponse()