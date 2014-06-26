package grocery

import (
	"testing"
)

func TestParseFoodStoresCSV(t *testing.T) {
	raw, err := ParseFoodStoresCSV("example_food_stores.csv")
	if err != nil {
		t.Fatalf("Error during ParseFoodStoresCSV: %s", err)
	}
	if len(raw) != 4 {
		t.Fatalf("Unexpected length of raw stores: %d", len(raw))
	}

	// Convert to stores
	stores, err := ConvertRawStores(raw)
	if err != nil {
		t.Fatalf("Error will converting raw stores: %s", err)
	}
	if len(stores) != 4 {
		t.Fatalf("Unexpected length of stores: %d", len(stores))
	}
}
