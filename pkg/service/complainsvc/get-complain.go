package complainsvc

import (
	"atommuse/backend/complain-services/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllComplaints retrieves all complaints
func (s *ComplainService) GetAllComplains(ctx context.Context) ([]model.Complain, error) {
	return s.repository.GetAllComplains(ctx)
}

func (s *ComplainService) GetComplainByExhibitionID(exhibitionID primitive.ObjectID) (*model.ResponseComplain, error) {
	return s.repository.GetComplainByExhibitionID(exhibitionID)
}

func (s *ComplainService) GetAllComplainsGroupByExhibitionName() ([]*model.ResponseComplain, error) {
	return s.repository.GetAllComplainsGroupByExhibitionName()
}
