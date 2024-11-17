package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nvi-identity-verification/business"
	"nvi-identity-verification/models"
)

type NviController struct {
	Business *business.NviBusiness
}

func NewNviController(business *business.NviBusiness) *NviController {
	return &NviController{Business: business}
}

func (ctrl *NviController) CheckTurkishNviHandler(c *gin.Context) {
	var request models.NviRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := ctrl.Business.CheckTurkishIdentity(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_valid": isValid})
}

func (ctrl *NviController) CheckForeignNviHandler(c *gin.Context) {
	var request models.ForeignNviRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := ctrl.Business.CheckForeignIdentity(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_valid": isValid})
}
