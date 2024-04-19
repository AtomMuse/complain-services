package complainrepo

import (
	"atommuse/backend/complain-services/pkg/model"
	"atommuse/backend/complain-services/pkg/utils"
	"context"
	"errors"
	"os"
	"strconv"

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

func (r *ComplainRepository) GetComplainByExhibitionID(exhibitionID primitive.ObjectID) (*model.ResponseComplain, error) {
	// Find the complains by ExhibitionID
	filter := bson.M{"exhibitionID": exhibitionID}
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each complain
	var complains []model.Complain
	for cursor.Next(context.Background()) {
		var complain model.Complain
		if err := cursor.Decode(&complain); err != nil {
			return nil, err
		}
		complains = append(complains, complain)
	}

	if len(complains) == 0 {
		return nil, errors.New("no complains found for the given ExhibitionID")
	}

	// Get the exhibition name from another database
	exhibition, err := r.getExhibitionByID(exhibitionID)
	if err != nil {
		return nil, err
	}

	// Construct the response
	response := &model.ResponseComplain{
		ExhibitionID:        exhibitionID,
		ExhibitionName:      exhibition.ExhibitionName,
		ExhibitionOrganizer: exhibition.ExhibitionOrganizer,
		ComplainMessage:     complains,
		BanRequests:         strconv.Itoa(len(complains)),
		Status:              model.StatusEnum(exhibition.Status),
		// You can fill in the rest of the fields accordingly
	}

	return response, nil
}

// getExhibitionByID retrieves the exhibition  by its ID from the exhibition collection
func (r *ComplainRepository) getExhibitionByID(exhibitionID primitive.ObjectID) (model.Exhibition, error) {
	// Connect to the MongoDB database
	client, err := utils.ConnectToMongoDB(os.Getenv("MONGO_URI"))
	if err != nil {
		return model.Exhibition{}, err
	}
	defer client.Disconnect(context.Background())

	// Specify the collection name
	exhibitionCollection := client.Database("atommuse").Collection("exhibitions")

	// Find the exhibition by ID
	var exhibition model.Exhibition

	filter := bson.M{"_id": exhibitionID}
	err = exhibitionCollection.FindOne(context.Background(), filter).Decode(&exhibition)
	if err != nil {
		return model.Exhibition{}, err
	}

	return exhibition, nil
}

func (r *ComplainRepository) GetAllComplainsGroupByExhibitionName() ([]*model.ResponseComplain, error) {
	// Retrieve all exhibition IDs
	exhibitionIDs, err := r.getAllExhibitionIDs()
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store detailed information about each exhibition
	var exhibitions []*model.ResponseComplain

	// Retrieve detailed information for each exhibition
	for _, id := range exhibitionIDs {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		// Retrieve complains for the current exhibition
		complains, err := r.GetComplainByExhibitionID(objectID)
		if err != nil {
			// If there are no complaints, continue to the next exhibition
			if err.Error() == "no complains found for the given ExhibitionID" {
				continue
			}
			return nil, err
		}
		exhibitions = append(exhibitions, complains)
	}

	return exhibitions, nil
}

func (r *ComplainRepository) getAllExhibitionIDs() ([]string, error) {
	// Connect to the MongoDB database
	client, err := utils.ConnectToMongoDB(os.Getenv("MONGO_URI"))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Specify the collection name
	exhibitionCollection := client.Database("atommuse").Collection("exhibitions")

	// Find all documents in the collection
	cursor, err := exhibitionCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and extract exhibition IDs
	var exhibitionIDs []string
	for cursor.Next(context.Background()) {
		var exhibition struct {
			ID primitive.ObjectID `bson:"_id,omitempty"`
		}
		if err := cursor.Decode(&exhibition); err != nil {
			return nil, err
		}
		exhibitionIDs = append(exhibitionIDs, exhibition.ID.Hex())
	}

	return exhibitionIDs, nil
}
