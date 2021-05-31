/*
リクエストで飛んでくるjson
*/
package request

type UpdateBalanceRequest struct {
	Amount        int    `json:"amount" binding:"required"`
	IdempotentKey string `json:"idempotentKey" binding:"required,min=36,max=36"`
}

type UpdateAllBalancesRequest struct {
	Amount        uint   `json:"amount"`
	IdempotentKey string `json:"idempotentKey" binding:"required,min=36,max=36"`
}
