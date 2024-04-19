package complainhanlder

import (
	"atommuse/backend/complain-services/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//	@Summary		Create a new complain
//	@Description	Create a new complain
//	@Tags			complain
//	@Security		BearerAuth
//	@ID				CreateComplain
//	@Accept			json
//	@Produce		json
//	@Param			requestExhibition	body	model.RequestCreateComplain	true	"Complain data to create"
//	@Success		201
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Failure		500
//	@Router			/api/complains [post]
func (h *ComplainHandler) CreateComplain(c *gin.Context) {
	// Get user information from request context
	userID, _ := c.Get("user_id")
	firstName, _ := c.Get("user_first_name")
	lastName, _ := c.Get("user_last_name")
	profileImage, _ := c.Get("user_image")

	// Convert user ID to primitive.ObjectID
	userIDObj, ok := userID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id is not a valid ObjectID"})
		return
	}

	// Parse request body to get complain data
	var reqBody model.RequestCreateComplain
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Populate User field with user information
	reqBody.User = model.User{
		ID:           userIDObj,
		FirstName:    firstName.(string),
		LastName:     lastName.(string),
		ProfileImage: profileImage.(string),
	}

	// Call Createcomplain method with userID
	complainID, err := h.service.CreateComplain(c, userIDObj, &reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"complainID": complainID})
}
