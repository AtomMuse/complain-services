package complainsvc

import (
	"atommuse/backend/complain-services/pkg/model"
	"context"
)

// GetAllComplaints retrieves all complaints
func (s *ComplainService) GetAllComplains(ctx context.Context) ([]model.Complain, error) {
	return s.repository.GetAllComplains(ctx)
}
