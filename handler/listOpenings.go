package handler

import (
	"fmt"
	"net/http"

	"github.com/HugoMarinn/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprint("unable to list openings"))
		return
	}

	sendSucess(ctx, "list-openings", openings)
}