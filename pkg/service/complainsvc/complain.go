package complainsvc

import "atommuse/backend/complain-services/pkg/repositorty/complainrepo"

// ComplaintService provides complaint-related operations
type ComplainService struct {
	repository complainrepo.ComplainRepository
}

// NewComplaintService creates a new ComplaintService
func NewComplainService(repository complainrepo.ComplainRepository) *ComplainService {
	return &ComplainService{
		repository: repository,
	}
}
