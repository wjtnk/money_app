/*
コントローラ
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"money-app/http/message"
	"money-app/http/request"
	"money-app/service/balance"
	"net/http"
)

type Controller struct {
	UpdateBalanceService     balance_service.UpdateBalanceService
	AddAllBalanceService     balance_service.AddAllBalanceService
	CreateSampleDataService  balance_service.CreateSampleDataService
	UpdateBalanceRequest     request.UpdateBalanceRequest
	UpdateAllBalancesRequest request.UpdateAllBalancesRequest
	ResponseMessage          message.ResponseMessage
}

func (bc Controller) UpdateBalance(c *gin.Context) {
	if err := c.ShouldBindJSON(&bc.UpdateBalanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, bc.ResponseMessage.GetErrorMessage(err.Error()))
		return
	}

	err := bc.UpdateBalanceService.Exec(c.Param("userId"), bc.UpdateBalanceRequest.Amount, bc.UpdateBalanceRequest.IdempotentKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, bc.ResponseMessage.GetErrorMessage(err.Error()))
		return
	}

	c.JSON(http.StatusOK, bc.ResponseMessage.GetSuccessMessage())
}

func (bc Controller) AddAllBalance(c *gin.Context) {
	if err := c.ShouldBindJSON(&bc.UpdateAllBalancesRequest); err != nil {
		c.JSON(http.StatusBadRequest, bc.ResponseMessage.GetErrorMessage(err.Error()))
		return
	}

	bc.AddAllBalanceService.Exec(bc.UpdateAllBalancesRequest.Amount, bc.UpdateAllBalancesRequest.IdempotentKey)

	c.JSON(http.StatusOK, bc.ResponseMessage.GetSuccessMessage())
}

func (bc Controller) CreateData(c *gin.Context) {
	bc.CreateSampleDataService.Exec()
	c.JSON(http.StatusOK, bc.ResponseMessage.GetSuccessMessage())
}
