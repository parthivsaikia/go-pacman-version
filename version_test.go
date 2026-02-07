package version

import (
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		v1   string
		v2   string
		want int
		name string
	}{
		// 1. Basic Equality
		{"1.0", "1.0", 0, "Identical versions"},
		{"1.0", "1.00", 0, "Leading zeros are equal"},

		// 2. Numeric Ordering
		{"1.0", "1.1", -1, "Simple numeric less than"},
		{"1.1", "1.0", 1, "Simple numeric greater than"},

		// 3. Epoch Handling
		{"1:1.0", "0:9.9", 1, "Epoch 1 beats Epoch 0"},

		// 4. Release Handling (Matches C logic: Skip if one is missing)
		{"1.0-1", "1.0-2", -1, "Same version, older release"},
		{"1.0", "1.0-1", 0, "One missing release = Skip comparison (Result 0)"}, // FIXED EXPECTATION

		// 5. RPM/Pacman Specifics
		{"1.0a", "1.0", -1, "Alpha is older than Empty"},
		{"1.0", "1-0", 1, "Dot separator vs Dash separator in Version (1.0 > 1)"}, // 1.0 parses as ver:1.0, 1-0 parses as ver:1 rel:0
		{"1.0", "1..0", -1, "Separator length matters (1 < 2)"},                   // FIXED EXPECTATION

		// 6. Edge Cases
		{"", "", 0, "Empty vs Empty"},
		{"1.0", "", 1, "Version vs Empty"},
		{"", "1.0", -1, "Empty vs Version"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare(tt.v1, tt.v2)
			if got != tt.want {
				t.Errorf("Compare(%q, %q) = %d; want %d", tt.v1, tt.v2, got, tt.want)
			}
		})
	}
}
