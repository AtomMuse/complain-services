package complainrepo

import (
	"atommuse/backend/complain-services/pkg/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllComplains retrieves all complains from the database
func (r *ComplainRepository) GetAllComplains(ctx context.Context) ([]model.Complain, error) {
	var complains []model.Complain
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &complains); err != nil {
		return nil, err
	}
	return complains, nil
}

func (r *ComplainRepository) GetComplainsByExhibitionID(ctx context.Context, exhibitionID primitive.ObjectID) ([]model.Complain, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{"exhibitionID": exhibitionID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var complains []model.Complain
	if err := cursor.All(ctx, &complains); err != nil {
		return nil, err
	}
	return complains, nil
}
