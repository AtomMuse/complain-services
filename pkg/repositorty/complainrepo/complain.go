package complainrepo

import "go.mongodb.org/mongo-driver/mongo"

// ComplaintRepository handles database operations related to complaints
type ComplainRepository struct {
	Collection *mongo.Collection
}

// NewComplaintRepository creates a new ComplaintRepository
func NewComplainRepository(collection *mongo.Collection) *ComplainRepository {
	return &ComplainRepository{Collection: collection}
}
