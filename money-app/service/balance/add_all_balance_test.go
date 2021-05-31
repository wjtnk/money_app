package balance_service

import (
	"money-app/db"
	"money-app/domain"
	"testing"
)

func TestAddAllBalanceService(t *testing.T) {
	db.Init()
	testDb := db.GetDB()
	defer db.DropAllTable()

	/**********
	 seedデータ準備
	**********/
	seedBalance1 := domain.Balance{Amount: 100, UserId: "1"}
	testDb.Create(&seedBalance1)

	seedBalance2 := domain.Balance{Amount: 200, UserId: "2"}
	testDb.Create(&seedBalance2)

	seedBalance3 := domain.Balance{Amount: 300, UserId: "3"}
	testDb.Create(&seedBalance3)

	/**********
	 テスト
	**********/
	a := AddAllBalanceService{}
	a.Exec(99, "sample-key")

	/**********
	 アサーション
	**********/
	balance1 := domain.Balance{}
	testDb.Where("user_id = ?", "1").First(&balance1)
	if balance1.Amount != 199 {
		t.Errorf("actual = %v, expect = %v", balance1.Amount, 199)
	}

	balance2 := domain.Balance{}
	testDb.Where("user_id = ?", "2").First(&balance2)
	if balance2.Amount != 299 {
		t.Errorf("actual = %v, expect = %v", balance2.Amount, 299)
	}

	balance3 := domain.Balance{}
	testDb.Where("user_id = ?", "3").First(&balance3)
	if balance3.Amount != 399 {
		t.Errorf("actual = %v, expect = %v", balance3.Amount, 399)
	}

	idempotentChecker := domain.IdempotentChecker{}
	testDb.Where("idempotent_key = ?", "sample-key").First(&idempotentChecker)
	if idempotentChecker.IdempotentKey != "sample-key" {
		t.Errorf("actual = %v, expect = %v", "sample-key", idempotentChecker.IdempotentKey)
	}
}

func TestAddAllBalanceService_冪等Keyがすでにあるときは加算されないこと(t *testing.T) {
	db.Init()
	testDb := db.GetDB()
	defer db.DropAllTable()

	/**********
	 seedデータ準備
	**********/
	seedBalance1 := domain.Balance{Amount: 100, UserId: "1"}
	testDb.Create(&seedBalance1)

	seedBalance2 := domain.Balance{Amount: 200, UserId: "2"}
	testDb.Create(&seedBalance2)

	seedBalance3 := domain.Balance{Amount: 300, UserId: "3"}
	testDb.Create(&seedBalance3)

	seedIdempotentChecker := domain.IdempotentChecker{IdempotentKey: "sample-key"}
	testDb.Create(&seedIdempotentChecker)

	/**********
	 テスト
	**********/
	a := AddAllBalanceService{}
	a.Exec(99, "sample-key")

	/**********
	 アサーション
	**********/
	balance1 := domain.Balance{}
	testDb.Where("user_id = ?", "1").First(&balance1)
	if balance1.Amount != 100 {
		t.Errorf("actual = %v, expect = %v", balance1.Amount, 100)
	}

	balance2 := domain.Balance{}
	testDb.Where("user_id = ?", "2").First(&balance2)
	if balance2.Amount != 200 {
		t.Errorf("actual = %v, expect = %v", balance2.Amount, 200)
	}

	balance3 := domain.Balance{}
	testDb.Where("user_id = ?", "3").First(&balance3)
	if balance3.Amount != 300 {
		t.Errorf("actual = %v, expect = %v", balance3.Amount, 300)
	}
}
