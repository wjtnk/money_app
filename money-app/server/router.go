/*
サービス起動処理とルーティング
*/
package server

import (
	"github.com/gin-gonic/gin"
	"money-app/http/controller"
)

func router() *gin.Engine {
	r := gin.Default()
	balanceCtrl := controller.Controller{}

	// 特定ユーザーの残高更新
	r.POST("/users/:userId/balance", balanceCtrl.UpdateBalance)

	// 残高を一括で更新
	r.POST("/balances", balanceCtrl.AddAllBalance)

	// データを作成
	r.POST("/sampleData", balanceCtrl.CreateData)

	return r
}
