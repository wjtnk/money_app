package domain

import "testing"

func Test_Update_Balance_Amountに加算されていること(t *testing.T) {
	balance := Balance{Amount: 0, UserId: "788b4d26-09c0-4390-99c4-0c7862f893f0"}
	balance.UpdateBalance(100)

	expect := 100
	if balance.Amount != expect {
		t.Errorf("actual = %v, expect = %v", balance.Amount, expect)
	}
}

func Test_Update_Balance_Amountから減算されていること(t *testing.T) {
	balance := Balance{Amount: 1000, UserId: "788b4d26-09c0-4390-99c4-0c7862f893f0"}
	balance.UpdateBalance(-900)

	expect := 100
	if balance.Amount != expect {
		t.Errorf("actual = %v, expect = %v", balance.Amount, expect)
	}
}

func Test_Update_Balance_Amountが0より下回らないこと(t *testing.T) {
	balance := Balance{Amount: 0, UserId: "788b4d26-09c0-4390-99c4-0c7862f893f0"}
	balance.UpdateBalance(-100)

	expect := 0
	if balance.Amount != expect {
		t.Errorf("actual = %v, expect = %v", balance.Amount, expect)
	}
}
