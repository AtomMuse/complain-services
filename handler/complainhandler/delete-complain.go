package complainhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		DeleteComplainByExhibitionID
// @Description	DeleteComplainByExhibitionID
// @Tags			Rejects
// @Security		BearerAuth
// @ID				DeleteComplainByExhibitionID
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"Exhibition ID"
// @Success		201
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router			/api-complains/complains/exhibitions/{id}/reject [delete]
func (h *ComplainHandler) DeleteComplainByExhibitionID(c *gin.Context) {
	// Extract exhibition ID from request parameters
	exhibitionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exhibition ID"})
		return
	}

	// Call the service method to delete complains by exhibition ID
	err = h.service.DeleteComplainByExhibitionID(exhibitionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete complains"})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Complains deleted successfully"})
}
