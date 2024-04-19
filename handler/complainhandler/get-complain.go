package complainhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary		GetAllComplains
// @Description	GetAllComplains
// @Tags			GetComplains
// @Security		BearerAuth
// @ID				GetAllComplains
// @Accept			json
// @Produce		json
// @Success		201
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router			/api-complains/complains/all [get]
func (h *ComplainHandler) GetAllComplains(c *gin.Context) {
	complains, err := h.service.GetAllComplains(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch complaints"})
		return
	}
	c.JSON(http.StatusOK, complains)
}

// @Summary		GetComplainByExhibitionID
// @Description	GetComplainByExhibitionID
// @Tags			GetComplains
// @Security		BearerAuth
// @ID				GetComplainByExhibitionID
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"Exhibition ID"
// @Success		201
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router			/api-complains/complains/exhibitions/{id} [get]
func (h *ComplainHandler) GetComplainByExhibitionID(c *gin.Context) {
	exhibitionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "Invalid comment ID"})
		return
	}

	complain, err := h.service.GetComplainByExhibitionID(exhibitionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errorMessage": "Complains not found"})
		return
	}

	c.JSON(http.StatusOK, complain)
}

// @Summary		GetComplainsGroupByExhibitionName
// @Description	GetComplainsGroupByExhibitionName
// @Tags			GetComplains
// @Security		BearerAuth
// @ID				GetComplainsGroupByExhibitionName
// @Accept			json
// @Produce		json
// @Success		201
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router			/api-complains/complains/exhibitions [get]
func (h *ComplainHandler) GetComplainsGroupByExhibitionName(c *gin.Context) {
	complainsMap, err := h.service.GetAllComplainsGroupByExhibitionName()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get complains grouped by exhibition name"})
		return
	}

	c.JSON(http.StatusOK, complainsMap)
}
