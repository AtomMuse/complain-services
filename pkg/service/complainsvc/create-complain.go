package complainsvc

import (
	"atommuse/backend/complain-services/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ComplainService) CreateComplain(ctx context.Context, userID primitive.ObjectID, complain *model.RequestCreateComplain) (*primitive.ObjectID, error) {
	return s.repository.CreateComplain(ctx, userID, complain)
}
