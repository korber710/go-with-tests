package pointers_and_errors

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepositsAndChecksBalance(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		w       Wallet
		deposit Bitcoin
	}{
		{Wallet{}, Bitcoin(10)},
		{Wallet{}, Bitcoin(5)},
	} {
		tc := tc
		t.Run(fmt.Sprintf("deposit %d and expect %d", tc.deposit, tc.deposit), func(t *testing.T) {
			t.Parallel()
			tc.w.Deposit(tc.deposit)
			assert.Equal(t, tc.w.Balance(), tc.deposit)
		})
	}
}

func TestWithdrawsAndChecksBalance(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		w        Wallet
		deposit  Bitcoin
		withdraw Bitcoin
		final    Bitcoin
	}{
		{Wallet{}, Bitcoin(10), Bitcoin(5), Bitcoin(5)},
		{Wallet{}, Bitcoin(5), Bitcoin(1), Bitcoin(4)},
	} {
		tc := tc
		t.Run(fmt.Sprintf("deposit %d, withdraw %d and expect %d", tc.deposit, tc.withdraw, tc.final), func(t *testing.T) {
			t.Parallel()
			tc.w.Deposit(tc.deposit)
			_ = tc.w.Withdraw(tc.withdraw)
			assert.Equal(t, tc.w.Balance(), tc.final)
		})
	}
}

func TestReportsErrorWhenInsufficientFunds(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		w        Wallet
		deposit  Bitcoin
		withdraw Bitcoin
	}{
		{Wallet{}, Bitcoin(10), Bitcoin(100)},
		{Wallet{}, Bitcoin(5), Bitcoin(6)},
	} {
		tc := tc
		t.Run(fmt.Sprintf("deposit %d, withdraw %d and expect error", tc.deposit, tc.withdraw), func(t *testing.T) {
			t.Parallel()
			tc.w.Deposit(tc.deposit)
			err := tc.w.Withdraw(tc.withdraw)
			assert.Error(t, err)
			assert.Equal(t, tc.w.Balance(), tc.deposit)
		})
	}
}
