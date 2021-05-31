package balance_service

import (
	"money-app/db"
	"money-app/domain"
	"testing"
)

func TestUpdateBalanceService(t *testing.T) {
	db.Init()
	testDb := db.GetDB()
	defer db.DropAllTable()

	/**********
	 seedデータ準備
	**********/
	seedBalance := domain.Balance{Amount: 100, UserId: "1"}
	testDb.Create(&seedBalance)

	/**********
	 テスト
	**********/
	u := UpdateBalanceService{}
	u.Exec("1", 200, "sample-key")

	/**********
	 アサーション
	**********/
	balance := domain.Balance{}
	testDb.Where("user_id = ?", "1").First(&balance)
	if balance.Amount != 300 {
		t.Errorf("actual = %v, expect = %v", balance.Amount, 300)
	}

	idempotentChecker := domain.IdempotentChecker{}
	testDb.Where("idempotent_key = ?", "sample-key").First(&idempotentChecker)
	if idempotentChecker.IdempotentKey != "sample-key" {
		t.Errorf("actual = %v, expect = %v", "sample-key", idempotentChecker.IdempotentKey)
	}
}

func TestUpdateBalanceService_冪等Keyがすでにあるときは更新されないこと(t *testing.T) {
	db.Init()
	testDb := db.GetDB()
	defer db.DropAllTable()

	/**********
	 seedデータ準備
	**********/
	seedBalance := domain.Balance{Amount: 100, UserId: "1"}
	testDb.Create(&seedBalance)

	seedIdempotentChecker := domain.IdempotentChecker{IdempotentKey: "sample-key"}
	testDb.Create(&seedIdempotentChecker)

	/**********
	 テスト
	**********/
	u := UpdateBalanceService{}
	u.Exec("1", 200, "sample-key")

	/**********
	 アサーション
	**********/
	balance := domain.Balance{}
	testDb.Where("user_id = ?", "1").First(&balance)
	if balance.Amount != 100 {
		t.Errorf("actual = %v, expect = %v", balance.Amount, 100)
	}
}

func TestUpdateBalanceService_更新前にエラーが発生した場合はロールバックされること(t *testing.T) {
	db.Init()
	testDb := db.GetDB()
	defer db.DropAllTable()

	/**********
	 seedデータ準備
	**********/
	seedBalance := domain.Balance{Amount: 100, UserId: "1"}
	testDb.Create(&seedBalance)

	/**********
	 テスト
	**********/
	u := UpdateBalanceService{}
	u.Exec("not-exist-user", 200, "new-sample")

	/**********
	 アサーション
	**********/
	// balancesとidempotent_checkersのどちらも保存、更新がされていないこと
	balance := domain.Balance{}
	testDb.Where("user_id = ?", "1").First(&balance)
	if balance.Amount != 100 {
		t.Errorf("actual = %v, expect = %v", balance.Amount, 100)
	}

	idempotentChecker := domain.IdempotentChecker{}
	testDb.Where("idempotent_key = ?", "new-sample").First(&idempotentChecker)
	if idempotentChecker.ID != 0 {
		t.Errorf("actual = %v, expect = %v", 0, idempotentChecker.ID)
	}
}
