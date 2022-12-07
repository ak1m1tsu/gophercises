package phone

import "testing"

type normalizeTestCase struct {
	input string
	want  string
}

func TestNormalize(t *testing.T) {
	testCases := []normalizeTestCase{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"(123)456-7892", "1234567892"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			actual := Normalize(testCase.input)
			if actual != testCase.want {
				t.Errorf("excepted %s, but got %s", testCase.want, actual)
			}
		})
	}
}
