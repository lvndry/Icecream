package supply_test

import (
	"testing"
)

func HasFlavor(t *testing.T) {
	if supply.HasFlavor("Vanilla") == false {
		t.Error("Expected True")
	}
}
