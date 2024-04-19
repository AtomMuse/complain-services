package complainrepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ComplainRepository) DeleteComplainByExhibitionID(exhibitionID primitive.ObjectID) error {
	// Define filter to delete complains by exhibition ID
	filter := bson.M{"exhibitionID": exhibitionID}

	// Delete documents matching the filter
	_, err := r.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
