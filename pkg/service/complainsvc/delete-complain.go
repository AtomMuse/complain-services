package complainsvc

import "go.mongodb.org/mongo-driver/bson/primitive"

// Service method to delete complains by exhibition ID
func (s *ComplainService) DeleteComplainByExhibitionID(exhibitionID primitive.ObjectID) error {
	// Call the repository method to delete complains by exhibition ID
	err := s.repository.DeleteComplainByExhibitionID(exhibitionID)
	if err != nil {
		return err
	}
	return nil
}
