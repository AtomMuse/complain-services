package complainrepo

import (
	"atommuse/backend/complain-services/pkg/model"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *ComplainRepository) CreateComplain(ctx context.Context, userID primitive.ObjectID, complain *model.RequestCreateComplain) (*primitive.ObjectID, error) {
	// Set UserID
	complain.User.ID = userID

	// Set Status
	complain.Status = "pending"

	// Set the Bangkok timezone offset to UTC+7
	const bangkokOffset = 7 * 60 * 60 // 7 hours in seconds

	// Get the current time in UTC
	currentTimeUTC := time.Now().UTC()

	// Add the Bangkok timezone offset to the current time
	currentTimeBangkok := currentTimeUTC.Add(time.Duration(bangkokOffset) * time.Second)

	// Set CreateDateAt in Bangkok timezone
	complain.CreateDateAt = primitive.NewDateTimeFromTime(currentTimeBangkok)

	// Insert the complain into the database
	result, err := r.Collection.InsertOne(ctx, complain)
	if err != nil {
		return nil, fmt.Errorf("failed to insert complain into the database: %w", err)
	}

	// Ensure InsertedID is not nil
	if result.InsertedID == nil {
		return nil, errors.New("failed to get inserted ID")
	}

	// Extract the generated ObjectID from the result
	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert inserted ID to ObjectID")
	}

	return &objectID, nil
}
