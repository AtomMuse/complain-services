package complainhandler

import "atommuse/backend/complain-services/pkg/service/complainsvc"

// ComplaintHandler handles complaint-related HTTP requests
type ComplainHandler struct {
	service complainsvc.ComplainService
}

// NewComplaintHandler creates a new ComplaintHandler
func NewComplainHandler(service complainsvc.ComplainService) *ComplainHandler {
	return &ComplainHandler{
		service: service,
	}
}
