package controllers

import (
	"century/models"
	"century/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BankController struct {
	service *services.BankService
}

func NewBankController(service *services.BankService) *BankController {
	return &BankController{service: service}
}

func (c *BankController) TransferHandler(ctx *gin.Context) {
	req := models.FundTransferRequest{}
	res := models.FundTransferResponse{
		Success: false,
	}
	if err := ctx.BindJSON(&req); err != nil {
		res.Error = "invalid request body"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.service.Transfer(req.From, req.To, req.Amount); err != nil {
		res.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.Success = true
	ctx.JSON(http.StatusOK, res)
}

func (c *BankController) BalanceHandler(ctx *gin.Context) {
	resp := models.BalanceResponse{
		Success: false,
	}
	accountID := ctx.Query("account_id")
	balance, err := c.service.GetBalance(accountID)
	if err != nil {
		resp.Error = err.Error()
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Success = true
	resp.Balance = balance
	ctx.JSON(http.StatusOK, resp)
	return
}
