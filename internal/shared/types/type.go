package types

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	Page  int
	Limit int
}

type SortOption struct {
	Field string
	Desc  bool
}

type Scope struct {
	University string
	Level      string
}

type ActionMeta struct {
	ActedBy string
	Reason  string
	Note    string
}

type PagedResult[T any] struct {
	Items []T
	Total uint64
	Page  int
	Limit int
}

type BoringResponse struct {
	Status int
	Length int
	Data   any
}

type UserFilter struct {
	Pagination
	Scope
	Search     string
	IsBanned   *bool
	Department string
	Faculty    string
	Sort       SortOption
	Level      string
}

type AdFilter struct {
	Pagination
	Scope
	IsActive *bool
	Slot     string
	Sort     SortOption
}

func (r BoringResponse) Write(w http.ResponseWriter) error {
	status := r.Status
	if status == 0 {
		status = http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(map[string]any{
		"status": r.Status,
		"length": r.Length,
		"data":   r.Data,
	})
}

func NewBoringResponse(status int, data any) BoringResponse {
	length := 0
	switch v := data.(type) {
	case nil:
		length = 0
	case string:
		length = len(v)
	case []any:
		length = len(v)
	}
	return BoringResponse{
		Status: status,
		Length: length,
		Data:   data,
	}
}
