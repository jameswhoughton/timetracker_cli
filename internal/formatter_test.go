package internal

import "testing"

func TestCaanRoundTimeToNearestGivenMinute(t *testing.T) {
	type TestCase struct {
		time     int
		roundBy  int
		expected int
	}

	testCases := []TestCase{
		{
			time:     25,
			roundBy:  4,
			expected: 24,
		},
		{
			time:     25,
			roundBy:  1,
			expected: 25,
		},
		{
			time:     25,
			roundBy:  15,
			expected: 30,
		},
	}

	for _, testCase := range testCases {
		t.Run("Rounding", func(t *testing.T) {
			result := Round(testCase.time, testCase.roundBy)

			if result != testCase.expected {
				t.Fatalf("Expected %d, got %d", testCase.expected, result)
			}
		})
	}
}

func TestRoundShouldReturnTheRoundByValueOrMore(t *testing.T) {
	expected := 15

	result := Round(1, 15)

	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
