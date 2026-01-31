package structs

import (
	"testing"
)

// The name of func NEED to start with Test and receive (t *testing.T)
func TestPaymentSlip_CalcData(t *testing.T) {
	// Scenario 1: Good return
	t.Run("Must to be calc 5%: ", func(t *testing.T) {
		slip := PaymentSlip{Id: 1, Amount: 100.0}

		// Call the method
		got, err := slip.CalcData(slip.Amount)
		want := 105.0 // 100 * 1.05

		// Verifications
		if err != nil {
			t.Errorf("Not wait error, but come: %v", err)
		}
		if got != want {
			t.Errorf("Wait %.2f, but come %.2f", want, got)
		}
	})

	// Scenario 2: Bad return
	t.Run("Need to return negative value", func(t *testing.T) {
		slip := PaymentSlip{Id: 1, Amount: -50.0}

		_, err := slip.CalcData(slip.Amount)

		// If err = nil, the test fails
		if err == nil {
			t.Error("Wait error, but not come")
		}
	})
}
