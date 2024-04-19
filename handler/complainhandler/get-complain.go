package complainhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		GetAllComplains
// @Description	GetAllComplains
// @Tags			Complains
// @Security		BearerAuth
// @ID				GetAllComplains
// @Accept			json
// @Produce		json
// @Success		201
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router			/api/complains [get]
func (h *ComplainHandler) GetAllComplains(c *gin.Context) {
	complains, err := h.service.GetAllComplains(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch complaints"})
		return
	}
	c.JSON(http.StatusOK, complains)
}
