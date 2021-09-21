package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/zharzhanov/region/models"
	"net/http"
)

func (h *Handler) addFeedback(c *gin.Context) {
	// get user id from middleware
	advertId := c.Query("advertId")

	var feedback models.Feedback
	if err := c.BindJSON(&feedback); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.AddFeedback(c.Request.Context(), feedback, advertId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}