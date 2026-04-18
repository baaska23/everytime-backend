package admin

import (
	"encoding/json"
	"net/http"
	"everytime-backend/internal/shared/types"
)

type ReportHandler struct {
	service ReportService
}

func NewReportHandler(service ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateReportRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		types.NewBoringResponse(http.StatusBadRequest, "invalid request body").Write(w)
		return
	}

	// basic validation before hitting service
	if req.PostID == "" || req.UserID == "" {
		types.NewBoringResponse(http.StatusBadRequest, "postId and userId are required").Write(w)
		return
	}

	report, err := h.service.Create(r.Context(), req)
	if err != nil {
		switch err.Error() {
		case "cannot report your own post":
			types.NewBoringResponse(http.StatusBadRequest, err.Error()).Write(w)
		case "already reported":
			types.NewBoringResponse(http.StatusConflict, err.Error()).Write(w)
		case "reason must be provided":
			types.NewBoringResponse(http.StatusBadRequest, err.Error()).Write(w)
		default:
			types.NewBoringResponse(http.StatusInternalServerError, "something went wrong").Write(w)
		}
		return
	}

	types.NewBoringResponse(http.StatusCreated, report).Write(w)
}